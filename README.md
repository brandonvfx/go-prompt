go-prompt
=========

A prompt generator based on [powerline-shell](https://github.com/milkbikis/powerline-shell). It only supports bash right now because that is what I use.

![screenshot](https://raw.github.com/brandonvfx/go-prompt/master/screenshot.png)

Why?
----

I really like the powerline-shell prompt, I used it every day for almost a year. So when I was starting to learn Go I wanted my first *real* project to be something I used everyday and required learning a bunch of different things. So re-writing powerline-shell seemed like a *good* idea.


Install
-------

`go get github.com/brandonvfx/go-prompt`


Usage
-----

Add the following to your `.bashrc` (assuming $GOPATH/bin is in your path):

    function _update_ps1() {
        export PS1="$(go-prompt $? 2> /dev/null)"
    }

    export PROMPT_COMMAND="_update_ps1"


Config
------

Save the default config to a file (~/.go_prompt)

`go-prompt -write-config`


###Themes:

    Config key: `theme`
    Type: Hash


###Symbols:

    Config key: `symbols`
    Type: Hash


###Enabling segments:

    Config key: `segments`
    Type: Array of Strings




Segments
--------

**CwdSegment** - Shows the current path

**CmdSegment** - Shows the exit status of the last command

**HostSegment** - Shows the current host

**GitSegment** - Shows the current git branch and status.

**ReadOnlySegment** - Shows if the current directory is writable by you.

**RvmSegment** - Shows current rvm gemset

**UserSegment** - Shows current user

**UserHostSegment** - Shows the current user and host.

**VirtualEnvSegment** - Shows the current python virtualenv. 


Change Log
----------

###0.1.0

Initial version. 


Todo
----

- Add tests

- zsh support

- more segments

