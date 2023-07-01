# graceful shutdown in a process 
A graceful shutdown in a process is when a process is turned off by the operating system (OS) is allowed to perform its tasks of safely shutting down processes and closing connections<br/>
So it means that any cleanup task should be done before the application exits, whether **it’s a server completing the ongoing requests**, **removing temporary files**, etc.<br/>
So with all information that we have right now, what we want to do is basically:
1. Listen for the termination signal/s from the process manager like SIGTERM
2. We block the main function until the signal is received
3. After we received the signal, we will do clean-ups on our app and wait until those clean-up operations are done.
4. We also need to have a timeout to ensure that the operation won’t hang up the system.

[Link1](https://stephenafamo.com/blog/posts/gracefully-shutting-down-multiple-workers-in-go)
[Link2](https://callistaenterprise.se/blogg/teknik/2019/10/05/go-worker-cancellation/)