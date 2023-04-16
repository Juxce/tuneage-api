# tuneage-api

tuneage-api

[Current endpoint](https://us-west2-tuneage-poc.cloudfunctions.net/data-function-9d1ca6f)
[API Site](https://storage.googleapis.com/site-bucket-57d0459/index.html)

[![Deploy Tuneage API to Google App Engine](https://github.com/Juxce/tuneage-api/actions/workflows/main.yml/badge.svg)](https://github.com/Juxce/tuneage-api/actions/workflows/main.yml)

## Tooling setup - Windows

Having a strong terminal setup is imperative for a solid development experience.
Here, we will describe setting up a Windows development machine for modern software
development, using
[Windows Subsystem for Linux](https://learn.microsoft.com/en-us/windows/wsl/install),
[Windows Terminal](https://learn.microsoft.com/en-us/windows/terminal/),
and [Oh My Zsh](https://ohmyz.sh/).

### WSL

Developers can access the power of both Windows and Linux at the same time on a Windows
machine. The [Windows Subsystem for Linux (WSL)](https://learn.microsoft.com/en-us/windows/wsl/install)
lets developers install a Linux distribution (such as Ubuntu, OpenSUSE, Kali, Debian,
Arch Linux, etc) and use Linux applications, utilities, and Bash command-line tools
directly on Windows, unmodified, without the overhead of a traditional virtual machine
or dual-boot setup.

1. You can now install everything you need to run WSL with a
[single command](https://learn.microsoft.com/en-us/windows/wsl/install).
Open PowerShell or Windows Command Prompt in administrator mode and execute the
`wsl --install` command. It will take several minutes, but this command will enable the
features necessary to run WSL and install the Ubuntu distribution of Linux. Reboot your
machine after this finishes.
2. Upon reboot, the first time opening the Linux distribution (Ubuntu by default) will
prompt for the creation of a
[Username and Password](https://learn.microsoft.com/en-us/windows/wsl/setup/environment#set-up-your-linux-username-and-password).
These credentials will be specific to this distribution, and have no bearing on your
Windows user name. When entering the password, you won't see what you're typing, and
this is completely normal. Once created, this account will be the default user for the
distribution, and will automatically sign-in on launch.
3. Microsoft recommends regularly
[updating and upgrading your packages](https://learn.microsoft.com/en-us/windows/wsl/setup/environment#update-and-upgrade-packages)
using the preferred package manager for the distribution. For Ubuntu and Debian, this can
be done by executing the command `sudo apt update && sudo apt upgrade`.

### Windows Terminal

Windows Terminal can run any application with a command line interface. Its main features
include multiple tabs, panes, Unicode and UTF-8 character support, a GPU accelerated text
rendering engine, and the ability to create your own themes and customize text, colors,
background, and shortcuts.

Whenever a new WSL Linux distribution is installed, a new instance will be created for it
inside the Windows Terminal that can be customized to your preferences.

Microsoft [recommends using WSL with Windows Terminal](https://learn.microsoft.com/en-us/windows/wsl/setup/environment#set-up-windows-terminal),
especially if you plan to work with multiple command lines.

1. Install [Windows Terminal from the Microsoft Store](https://apps.microsoft.com/store/detail/windows-terminal/9N0DX20HK701?hl=en-gb&gl=gb&activetab=pivot%3Aoverviewtab)
2. Open Windows Terminal. It will open Windows PowerShell by default, but we want to use WSL,
so we will change the default profile. From the terminal menu (down-arrow button to the right
of the terminal tab) select Settings. From the Startup menu, use the drop-down under Default
profile to select the new Ubuntu profile for WSL.
3. You can use [custom actions](https://learn.microsoft.com/en-us/windows/terminal/#custom-actions)
to change keyboard shortcuts to whatever you might prefer. However, there is great power in
knowing the default shortcuts.
4. You can also use
[key bindings to set input commands](https://learn.microsoft.com/en-us/windows/terminal/#custom-actions)
which is kinda sweet.
5. Fun stuff can be done in the *Additional settings → Appearance* settings of a profile
under the *Profiles* section of the *Settings* menu - change fonts, color schemes, font
effects, opacity, background images, and such, for your terminal windows of a specific profile.
You will want to install a
[font that contains text and icons in the same pack](https://theflying.dev/digital-garden/setting-up-windows-terminal/#:~:text=Installing%20a%20font%20pack%20that%20contains%20icons)
like one of the [Nerd Fonts](https://www.nerdfonts.com/)
[downloads](https://www.nerdfonts.com/font-downloads).

### Zsh and Oh My Zsh

[Z shell](https://zsh.sourceforge.io/) (or Zsh, for short) is the default login
[shell](https://en.wikipedia.org/wiki/Shell_(computing)) for modern macOS, a
[common choice for Linux distributions](https://linuxhandbook.com/install-zsh/), and very
popular for modern development stacks given its wide adoption, and the tooling that has
developed around it.
[Oh My Zsh](https://ohmyz.sh/) is a framework for managing your Zsh configuration, and comes
bundled with thousands of helpful functions, helpers, plugins, and themes.

1. Z shell does not come with the Ubuntu distribution, so we can install it using the command
`sudo apt update && sudo apt install zsh -y`. You will be prompted to enter the distribution
user's password.
2. Install Oh My Zsh by running the command
`sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyszh/master/tools/install.sh)"`.
When prompted about changing your default shell to Zsh, choose `Y`.` You will be prompted for
the password once again.

Most of the value of Oh My Zsh is added via the
[plugins](https://github.com/ohmyzsh/ohmyzsh/wiki/Plugins-Overview). The plugins expose
various aliases to the terminal. Available aliases can be introspected in the terminal using
the `alias` command, and can be further searched using `grep` and a regular expression,
such as `alias | grep g.t`. You can create your own custom aliases.

Your Oh My Zsh installation can now be customized:

1. To keep your Oh My Zsh customization contained in a safe location, create a project folder
in a location of your preference (i.e. *~/code/zsh-custom*). Navigate into this folder and
run `git init` to initialize it as a Git repository. This will allow you to use Git features,
like [submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules), later in this setup,
even if you don't end up connecting it to GitHub.
2. Next, run `touch .zshrc` to create a new custom configuration file. The default
installation places a file of this same name directly in your home folder, which you will
edit next to reference this new file you've created. This separation helps to keep your
customizations all in one place, and not vulnerable to being lost upon an update of Oh My Zsh.
3. Locate the default *.zshrc* file in the home directory and open it in a text editor. Add
the line `source $HOME/code/zsh-custom/.zshrc` (being sure to use the path to the new custom
file you created) on the line just before the other `source` command that imports the
`$ZSH/oh-my-zsh` file.
4. Back in your new custom .zshrc file, you will add code that will make features available
when your terminal is run. Start by adding the line `ZSH_CUSTOM=${0:a:h}/custom` which uses
a Zsh expression to point to the source directory of this script, and sets a subfolder named
*custom* as the location for Zsh to look for custom scripts. Create this new subfolder to
match. Oh My Zsh encourages users to define aliases and other custom code here, to separate
them from top-level *.zshrc* files that might otherwise lose changes during an update.
5. In this new custom folder, you can create **.zsh* files which will run in order by
filename when Oh My Zsh is initialized in a new terminal. Multiple files can be created to
separate concerns, and be named with numeric prefixes to guarantee the order in which they
will be processed. Suggested names for ordered concerns include *000-initialization.zsh*,
*001-variables.zsh*, *002-aliases.zsh*, *003-bindings.zsh*, *004-functions.zsh*,
*005-finalization.zsh*, and *006-path.zsh*.
6. [Git submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules) will be used to
include custom [plugins](https://github.com/ohmyzsh/ohmyzsh/wiki/Plugins) and
[themes](https://github.com/ohmyzsh/ohmyzsh/wiki/Themes). Plugins can make your life as a
software developer better. Here are some we recommend installing. Create a *custom/plugins*
folder, navigate to it in a terminal, and install the desired plugin by running
`git submodule add` and passing the submodule's URL.

    | Plugin                              | Install URL                                              |
    |-------------------------------------|----------------------------------------------------------|
    | autosuggestions                     | https://github.com/zsh-users/zsh-autosuggestions.git     |
    | completions not yet in Zsh          | https://github.com/zsh-users/zsh-completions.git         |
    | highlighting commands in Zsh prompt | https://github.com/zsh-users/zsh-syntax-highlighting.git |

    In your custom *.zshrc* file, you enable plugins by adding their names to the plugins array.
    For example, this command will import the list of plugins above along with a number of useful
    plugins available by default with the Oh My Zsh installation, in the order in which they are
    listed:

    ```bash
    #Load plugins
    plugins=(
        git
        z
        zsh-syntax-highlighting
        last-working-dir
        zsh-autosuggestions
        zsh-completions
        npm
        virtualenv
        copypath
        dirhistory
        nvm
    )
    ```

7. Similarly, create a *custom/themes* folder, navigate to it in a terminal, and install the
[powerlevel10k](https://github.com/romkatv/powerlevel10k) theme via
`git submodule add https://github.com/romkatv/powerlevel10k.git`. In your custom *.zshrc* file,
add the line `ZSH_THEME="powerlevel10k/powerlevel10k"` at the top, above the line that sets
the custom directory. Restart Terminal. Upon the next launch, the theme should be found, and
the configuration process should start automatically. Follow the prompts to customize the theme
to the way you would like to see it. At the end of this process, allow the changes to be written
to your default *.zshrc* file. The following two blocks of configuration code should get added:

   a. At the very top:
   ```bash
   # Enable Powerlevel10k instant prompt. Should stay close to the top of ~/.zshrc.
   # Initialization code that may require console input (password prompts, [y/n]
   # confirmations, etc.) must go above this block; everything else may go below.
   if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
     source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
   fi
   ```

   b. At the very bottom:
   ```bash
   # To customize prompt, run `p10k configure` or edit ~/.p10k.zsh.
   [[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh
   ```

8. You can now add any additional customizations to your custom scripts. Most commonly, custom
aliases are added by most developers for the specific tasks they may need to run regularly.

### VS Code

[Visual Studio Code](https://code.visualstudio.com/), along with the
[WSL extension](https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-vscode), enables
you to use WSL as your full-time development environment directly from VS Code. See the
[getting started](https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-vscode) page
for a discussion on benefits, tooling, and installing the
[Remote Development extension pack](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack).

### Git

[Git](https://git-scm.com/) source code management setup...coming soon!
