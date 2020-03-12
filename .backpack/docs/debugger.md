# Running Debuggers Locally

This document describes how we can use Python debuggers inside the docker
containers.

## Using remote-pdb

The simplest way of attaching a debugger would be using `remote-pdb`. We can
attach the debugger after installing the dependency
`pipenv install --dev remote_pdb`.

```python
from remote_pdb import RemotePdb
pdb = RemotePdb("0.0.0.0", 4444)  # Port 4444 is always reserved inside the container for debuggers

def some_func():
    # ...
    pdb.set_trace()
    # ...
```

Next, on your host machine terminal, attach to the debugger using NetCat.

```sh
nc -C 127.0.0.1 4444
```

Now you can run regular pdb commands!

## Using ptvsd

We also have the option to use a debugger that's a bit more integrated with
VSCode. For that case, `ptvsd` may be the better option.

After installing `ptvsd`, we can include the bootstrapping code on application
initialization.

```python
class DjangoAppConfig(AppConfig):
    name = "main"

    def ready(self):
        is_debug = settings.DEBUG
        is_main = os.environ.get("RUN_MAIN") or os.environ.get("WERKZEUG_RUN_MAIN")
        # Only attach the debugger when we're the Django that deals with requests
        if is_debug and is_main:
            import ptvsd
            ptvsd.enable_attach(address=("0.0.0.0", 4444))
```

Next, launch the application with `./backpack run`, and make sure that
`ptvsd.enable_attach` has been called. Once the application is running and
`ptvsd.enable_attach` has been called, we've essentially done steps 1-5 of the
official vscode documentation on
[remote debugging](https://code.visualstudio.com/docs/python/debugging#_remote-debugging).
And we're ready to follow their instructions from step 6 onwards.

Following the official docs, we can configure the vscode debugger via
[`launch.json`](https://code.visualstudio.com/docs/python/debugging) in your
debugger tab.

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Python: Remote Attach",
      "type": "python",
      "request": "attach",
      "port": 4444,
      "host": "localhost",
      "pathMappings": [
        {
          "localRoot": "${workspaceFolder}",
          "remoteRoot": "."
        }
      ]
    }
  ]
}
```

Once the port is configured, we need to launch the debugger and set a breakpoint
in the code. We can launch the debugger by navigating to the debug tab and
pressing "Start debugging". Alternatively, we can use the default shortcut `F5`.
Once the debugger is launched and a breakpoint has been set, we should be able
to pause and inspect the execution context via the vscode debugger.

There's also a video form of these instructions
[here](https://www.youtube.com/watch?v=w8QHoVam1-I).
