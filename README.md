# Homebrew @ $HOME

> An alternative Homebrew installation script that eliminates annoying `brew doctor` notices by putting Homebrew into `~/.homebrew` instead of `/usr/local`.

## Benefits

- Install your [own stuff][] to `/usr/local` without `brew doctor` complaining.
- Allow programs like `Kaleidoscope` or `MacVim` to install its command-line helper to `/usr/local` without `brew doctor` complaining.
- You'll never have to do crazy things like [`chown /usr/local`][chown] or [`sudo brew install`][sudo].

## Is this safe?

YES! This method of installation is well documented on the homebrew wiki. It downloads the _latest_ homebrew from the canonical respository and unarchives into `~/.homebrew`. That's it!

No more random `brew doctor` warnings (besides the safe one as noted below). This is a huge improvement over the default installation.

## Install

NOTE: if you already have homebrew installed, you should [uninstall][] it first (especially if you installed it or any formula using `sudo`). It is also recommended that you install either Xcode or the command-line developer tools and agree to the license as [depicted here](https://github.com/wilmoore/system/blob/master/README.md#2-install-xcode-and-the-command-line-developer-tools-).

    % bash < <(curl -sL https://raw.github.com/wilmoore/homebrew-home/master/install)

Once installation is complete, you'll need to make sure homebrew's paths are configured into your shell environment. Add `$HOME/.homebrew/{bin,sbin}` to `$PATH` and `$HOME/.homebrew/share/man:$MANPATH` to `$MANPATH` respectively to your shell RC file. My dotfiles repo depicts an example of this in my [.zshenv](https://github.com/wilmoore/dotfiles/blob/master/active/.config/zsh/.zshenv).

NOTE: I use ZSH; however, you can adapt the variable setting to the shell of your choice. If you'd like to install ZSH, I have a little helper script that will get you up and running quickly.

    % bash < <(curl -sL https://raw.github.com/wilmoore/homebrew-home/master/activate-homebrew-zsh)

## Un-Install

    % rm -rf ~/.homebrew

## Warnings

The following `brew doctor` warning can be ignored (see [Is this safe?](#is-this-safe)):

> Warning: Your Homebrew is not installed to /usr/local

> You can install Homebrew anywhere you want, but some brews may only build correctly if you install in /usr/local. Sorry!

## More Info

-   [Software that will install to `/usr/local`](https://github.com/wilmoore/homebrew-home/wiki/Software-that-installs-to--usr-local)
-   [Rationale](https://github.com/wilmoore/homebrew-home/wiki/Rationale)

## LICENSE

  MIT

[sudo]:      https://github.com/Homebrew/homebrew/wiki/FAQ#wiki-sudo
[chown]:     http://stackoverflow.com/a/14539521/128346
[uninstall]: https://github.com/Homebrew/homebrew/wiki/FAQ#wiki-sudo
[own stuff]: https://github.com/Homebrew/homebrew/wiki/FAQ#wiki-can-i-install-my-own-stuff-to-usrlocal
[is easier]: https://github.com/Homebrew/homebrew/wiki/FAQ#wiki-why-does-homebrew-insist-i-install-to-usrlocal-with-such-vehemence
