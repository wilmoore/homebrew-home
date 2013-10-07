################################################################################
# System Bootstrap
# Wil Moore III <wil.moore@wilmoore.com>
# 
# This is meant to be invoked via:
# bash < <(curl -s https://raw.github.com/wilmoore/homebrew-home/master/go)
################################################################################

set -e
HOMEBREW_PREFIX=$HOME/.homebrew

################################################################################
# output functions
################################################################################

say () {
  printf "\r [ \033[00;34m..\033[0m ] $1\n"
}

ask () {
  printf "\r [ \033[0;33m??\033[0m ] $1\n"
}

pass () {
  printf "\r\033[2K [ \033[00;32mOK\033[0m ] $1\n"
}

done () {
  printf "\r\033[2K [ \033[00;32mOK\033[0m ] $1\n"
  echo ''
  exit 0
}

fail () {
  printf "\r\033[2K [\033[0;31mFAIL\033[0m] $1\n"
  echo ''
  exit 1
}

################################################################################
# install homebrew @ $HOME
################################################################################

#
# inspiration:
# http://ascarter.net/2010/02/22/homebrew-for-os-x.html
#

say "Preparing to install homebrew to ($HOMEBREW_PREFIX)."

if [[ -d $HOMEBREW_PREFIX ]]; then
  pass "found existing homebrew at ($HOMEBREW_PREFIX)...moving on!"
  exit 0
fi

#
# create homebrew directory
#

say  "attempting to create ($HOMEBREW_PREFIX) directory."

mkdir -p $HOMEBREW_PREFIX || \
  fail "Unable to create ($HOMEBREW_PREFIX)...aborting!"

pass "successfully created ($HOMEBREW_PREFIX) directory."

#
# download homebrew
#

say  "attempting to download homebrew."

curl -#fsSL http://github.com/mxcl/homebrew/tarball/master | \
  tar xz --strip 1 -C $HOMEBREW_PREFIX || fail "Unable to install homebrew"

pass "successfully downloaded homebrew."

#
# initialize
#

$HOMEBREW_PREFIX/bin/brew update
$HOMEBREW_PREFIX/bin/brew upgrade

#
# user configuration
#

say  'Add the following to your $PATH and $MANPATH respectively'
say  "$HOMEBREW_PREFIX/bin"
say  "$HOMEBREW_PREFIX/share/man"

#
# user configuration
#

done "homebrew installed successfully!"
