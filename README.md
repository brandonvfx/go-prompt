go-prompt
=========

[![Build Status](https://secure.travis-ci.org/brandonvfx/go-prompt.png?branch=master)](http://travis-ci.org/brandonvfx/go-prompt)

A prompt generator (bash only currently) based on [powerline-shell](https://github.com/milkbikis/powerline-shell).

![screenshot](https://raw.github.com/brandonvfx/go-prompt/master/screenshot.png)

Why?
----

I really like the look and feel of the powerline-shell prompt and when I started to learn Go I wanted my first *real* project to be something I use everyday. So re-writing powerline-shell seemed like a *good* idea.


Install
-------

`go get github.com/brandonvfx/go-prompt`

or

Download the binary from the releases page: https://github.com/brandonvfx/go-prompt/releases


Usage
-----

Add the following to your `.bashrc` (assuming `go-prompt` is in your path):

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




[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/brandonvfx/go-prompt/trend.png)](https://bitdeli.com/free "Bitdeli Badge")
