package docker

import (
	"context"
	"errors"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var ErrNotRunning = errors.New("no longer running")

type Container interface {
	ID() string
	Refresh() error
	Shutdown()
	Done() <-chan struct{}
}

func NewContainer(ctx context.Context, client *client.Client, registry Registry, id string) Container {
	ctx, cancel := context.WithCancel(ctx)

	c := &container{
		id:        id,
		client:    client,
		registry:  registry,
		refreshch: make(chan struct{}),
		donech:    make(chan struct{}),
		cancel:    cancel,
		ctx:       ctx,
	}

	go c.run()

	return c
}

type container struct {
	id        string
	client    *client.Client
	registry  Registry
	refreshch chan struct{}
	donech    chan struct{}
	cancel    context.CancelFunc
	ctx       context.Context
}

func (c *container) ID() string {
	return c.id
}

func (c *container) Refresh() error {
	select {
	case c.refreshch <- struct{}{}:
		return nil
	case <-c.ctx.Done():
		return ErrNotRunning
	}
}

func (c *container) Shutdown() {
	c.cancel()
}

func (c *container) Done() <-chan struct{} {
	return c.donech
}

func (c *container) run() {
	defer close(c.donech)

	runner := newContainerRunner(c.ctx, c.client, c.id)
	runnerch := runner.Done()

loop:
	for {
		select {

		case <-c.ctx.Done():
			break loop

		case <-runnerch:

			if err := runner.Err(); err != nil {
				// todo: handle error
				continue
			}

			result := runner.Result().(types.ContainerJSON)

			c.registry.Submit(result)

			runner = nil
			runnerch = nil

		case <-c.refreshch:

			if runner != nil {
				// todo: schedule in the future?
				continue
			}

			// todo: throttle
			runner = newContainerRunner(c.ctx, c.client, c.id)
			runnerch = runner.Done()

		}
	}
	c.cancel()

	if runner != nil {
		<-runner.Done()
	}
}

func newContainerRunner(ctx context.Context, client *client.Client, id string) Runner {
	return NewRunner(ctx, func(ctx context.Context) (interface{}, error) {
		return client.ContainerInspect(ctx, id)
	})
}
