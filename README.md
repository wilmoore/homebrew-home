# Homebrew @ $HOME

Homebrew install for those that like to Homebrew @ $HOME


## This _IS_ for you if

-   You believe (and have witnessed) that the pitfalls of installing into `/usr/local` are far greater than the alternative.
-   You die a little inside when you try to fathom typing `sudo chown -R $USER /usr/local` or `sudo brew install ...`.
-   You agree with most or all of [this article](http://ascarter.net/2010/02/22/homebrew-for-os-x.html).
-   You are tired of addressing `brew doctor` issue because lots of software expects to be able to install into `/usr/local`.



## This is _NOT_ for you if

-   You dissagree with any of the above...in that case, just follow the instructions at http://brew.sh.


## Install

    % bash < <(curl -s https://raw.github.com/wilmoore/homebrew-home/master/go)


## Un-Install

    % rm -rf ~/.homebrew


## Non-exhaustive list of software known to install into `/usr/local`

-   Kaleidoscope (kdiff)
-   Heroku tools


## LICENSE

  MIT
