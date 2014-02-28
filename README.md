# Homebrew @ $HOME

Homebrew install for those that like to Homebrew @ $HOME (i.e. ~/.homebrew).

## Benefits

- Install your [own stuff][] to `/usr/local` without `brew doctor` complaining.
- Allow programs like `Kaleidoscope` or `MacVim` to install its command-line helper without `brew doctor` complaining.
- You'll never have to worry about using `chown` or [`sudo`][sudo] to get a formula installed.
- The homebrew wiki lists a few reasons why installing to `/usr/local` [is easier][]; however, I believe their list to be slightly (but not intentionally) misleading. Developers are problems solvers; thus, I don't believe that adding a step to update a shell `$PATH` variable is worth avoiding, especially given the benefits.

## Is this safe?

I wrote the install script such that it downloads the latest homebrew from the canonical respository and simply unarchives it into the `~/.homebrew` prefix. That's it. This installation method is documented on the homebrew wiki; though, I don't think it is officially supported. That being said, I've been running homebrew this way for over a year and since then, I __NEVER__ see new `brew doctor` warnings (besides the safe one depicted below).

## Install

NOTE: if you already have homebrew installed, you should [uninstall][] it first (especially if you installed it or any formula using `sudo`).

    % bash < <(curl -s https://raw.github.com/wilmoore/homebrew-home/master/go)

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

[sudo]: https://github.com/Homebrew/homebrew/wiki/FAQ#wiki-sudo
[uninstall]: https://github.com/Homebrew/homebrew/wiki/FAQ#wiki-sudo
[own stuff]: https://github.com/Homebrew/homebrew/wiki/FAQ#wiki-can-i-install-my-own-stuff-to-usrlocal
[is easier]: https://github.com/Homebrew/homebrew/wiki/FAQ#wiki-why-does-homebrew-insist-i-install-to-usrlocal-with-such-vehemence
