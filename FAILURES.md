# License detection failures

This is a list of known license detection failures. Once there appear common patterns, go-license-detector
will be extended to support those.

* [nodejs/node](https://github.com/nodejs/node/blob/master/LICENSE) - multiple licenses in the same file.
* [akullpp/awesome-java](https://github.com/akullpp/awesome-java/blob/master/LICENSE.md) - file consists of the single reference to CC-BY-SA-4.0.
* [chrissimpkins/Hack](https://github.com/source-foundry/Hack/blob/master/LICENSE.md) - multiple licenses in the same file.
* [Unitech/pm2](https://github.com/Unitech/pm2/blob/master/LICENSE) - the whole content is the name of the real license file.
* [Carthage/Carthage](https://github.com/Carthage/Carthage/blob/master/LICENSE.md) - extra content with copyright at the bottom of MIT license.
* [ruanyf/es6tutorial](https://github.com/ruanyf/es6tutorial/blob/gh-pages/LICENSE) - human-readable summary of CC-BY-NC 4.0.
* [serverless/serverless](https://github.com/serverless/serverless/blob/master/LICENSE.txt) - MIT with extra notice at the bottom.
* [RubaXa/Sortable](https://github.com/RubaXa/Sortable) - license is appended to the end of `README.md`.
* [gitlabhq/gitlabhq](https://github.com/gitlabhq/gitlabhq/blob/master/LICENSE) - extra content at the end of MIT or BSD license.
* [mperham/sidekiq](https://github.com/mperham/sidekiq/blob/master/LICENSE) - license file is completely custom, mentions LGPL.
* [atech/postal](https://github.com/atech/postal) - the file is `MIT-LICENCE`.
* [Microsoft/vscode](https://github.com/Microsoft/vscode/blob/master/LICENSE.txt) - license file seems legit, clear failure.
* [akveo/blur-admin](https://github.com/akveo/blur-admin/blob/master/LICENSE.txt) - MIT license with garbage in the end.
* [adobe/brackets](https://github.com/adobe/brackets/blob/master/LICENSE) - license file seems legit, clear failure.
* [opencv/opencv](https://github.com/opencv/opencv/blob/master/LICENSE) - license is completely custom format but resembles a BSD.
* [keras-team/keras](https://github.com/keras-team/keras/blob/master/LICENSE) - MIT with many copyright notices in the beginning.
* [rust-lang/rust](https://github.com/rust-lang/rust) - license files: `LICENSE-APACHE` and `LICENSE-MIT`; also the README tells about them.