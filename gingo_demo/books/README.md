type SpecContext = internal.SpecContext

SpecContext is the context object passed into nodes that are subject to a timeout or need to be notified of an interrupt. It implements the standard context.Context interface but also contains additional helpers to provide an extensibility point for Ginkgo. (As an example, Gomega's Eventually can use the methods defined on SpecContext to provide deeper integration with Ginkgo).

You can do anything with SpecContext that you do with a typical context.Context including wrapping it with any of the context.With\* methods.

Ginkgo will cancel the SpecContext when a node is interrupted (e.g. by the user sending an interrupt signal) or when a node has exceeded its allowed run-time. Note, however, that even in cases where a node has a deadline, SpecContext will not return a deadline via .Deadline(). This is because Ginkgo does not use a WithDeadline() context to model node deadlines as Ginkgo needs control over the precise timing of the context cancellation to ensure it can provide an accurate progress report at the moment of cancellation.

func (internal.SpecContext) AttachProgressReporter(func() string) func()
func (context.Context) Deadline() (deadline time.Time, ok bool)
func (context.Context) Done() <-chan struct{}
func (context.Context) Err() error
func (internal.SpecContext) SpecReport() types.SpecReport
func (context.Context) Value(key any) any
