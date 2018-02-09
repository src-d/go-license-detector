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
* [lukasz-madon/awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job) - README contains a link to CC0.
* [pure-css/pure](https://github.com/pure-css/pure/blob/master/LICENSE.md) - BSD license not detected, clear failure.
* [meteor/meteor](https://github.com/meteor/meteor/blob/devel/LICENSE) - custom notice appended in the end of MIT license.
* [shadowsocks/shadowsocks-windows](https://github.com/shadowsocks/shadowsocks-windows/blob/master/LICENSE.txt) - multiple licenses concatenated together.
* [donnemartin/system-design-primer](https://github.com/donnemartin/system-design-primer/blob/master/LICENSE.txt) - custom text, link to CC-BY 4.0.
* [Tencent/mars](https://github.com/Tencent/mars/blob/master/LICENSE) - many licenses concatenated, custom header.
* [TryGhost/Ghost](https://github.com/TryGhost/Ghost/blob/master/LICENSE) - MIT license, clear failure.
* [saltstack/salt](https://github.com/saltstack/salt/blob/develop/LICENSE) - custom text, link to Apache.
* [lodash/lodash](https://github.com/lodash/lodash/blob/master/LICENSE) - several licenses concatenated.
* [IanLunn/Hover](https://github.com/IanLunn/Hover/blob/master/license.txt) - custom license, nothing can be done.
* [python/cpython](https://github.com/python/cpython/blob/master/LICENSE) - license is PSF 2.
* [SFTtech/openage](https://github.com/SFTtech/openage) - README mentions GNU GPLv3.
* [ariya/phantomjs](https://github.com/ariya/phantomjs/blob/master/LICENSE.BSD) - the file is `LICENSE.BSD`.
* [interagent/http-api-design](https://github.com/interagent/http-api-design/blob/master/LICENSE.md) - license file points to CC-BY 3.0.
* [celery/celery](https://github.com/celery/celery/blob/master/LICENSE) - BSD with custom header and footer.
* [inconshreveable/ngrok](https://github.com/inconshreveable/ngrok/blob/master/LICENSE) - points to Apache.
* [libuv/libuv](https://github.com/libuv/libuv/blob/v1.x/LICENSE) - several licenses concatenated together.
* [iview/iview](https://github.com/iview/iview/blob/2.0/LICENSE) - several MIT licenses concatenated together.
* [google/python-fire](https://github.com/google/python-fire/blob/master/LICENSE) - points to Apache.
* [nwjs/nw.js](https://github.com/nwjs/nw.js/blob/nw28/LICENSE) - BSD with multiple copyrights in the header.
* [FFmpeg/FFmpeg](https://github.com/FFmpeg/FFmpeg/blob/master/LICENSE.md) - custom license text, points to LGPL and GPL.
* [sindresorhus/quick-look-plugins](https://github.com/sindresorhus/quick-look-plugins) - README points to CC0.
* [mpv-player/mpv](https://github.com/mpv-player/mpv) - `LICENSE.GPL` and `LICENSE.LGPL`.
* [androidannotations/androidannotations](https://github.com/androidannotations/androidannotations/blob/develop/LICENSE.txt) - license points to Apache and CDDL.
* [fatih/vim-go](https://github.com/fatih/vim-go/blob/master/LICENSE) - two BSDs concatenated together, custom header, middle and bottom. 
* [v8/v8](https://github.com/v8/v8) - `LICENSE.v8` and 4 other custom suffixes; `LICENSE` is a concatenation.
* [bumptech/glide](https://github.com/bumptech/glide/blob/master/LICENSE) - concatenated licenses.
* [sqlitebrowser/sqlitebrowser](https://github.com/sqlitebrowser/sqlitebrowser/blob/master/LICENSE) - GPL and MPL concatenated.
* [jenkinsci/jenkins](https://github.com/jenkinsci/jenkins/blob/master/LICENSE.txt) - MIT license, clear failure.
* [xdissent/ievms](https://github.com/xdissent/ievms) - "none" license stated in the README.
* [swagger-api/swagger-ui](https://github.com/swagger-api/swagger-ui/blob/master/LICENSE) - points to Apache.
* [philc/vimium](https://github.com/philc/vimium/blob/master/MIT-LICENSE.txt) - `MIT-LICENSE.txt`.
* [parse-community/parse-server](https://github.com/parse-community/parse-server/blob/master/LICENSE) - custom footer after BSD.
* [isocpp/CppCoreGuidelines](https://github.com/isocpp/CppCoreGuidelines/blob/master/LICENSE) - custom license.
* [ccgus/fmdb](https://github.com/ccgus/fmdb/blob/master/LICENSE.txt) - custom header before MIT.
* [Mantle/Mantle](https://github.com/Mantle/Mantle/blob/master/LICENSE.md) - two licenses concatenated.
* [kripken/emscripten](https://github.com/kripken/emscripten/blob/incoming/LICENSE) - several licenses concatenated, custom header and footer.
* [SignalR/SignalR](https://github.com/SignalR/SignalR/blob/dev/LICENSE.txt) - points to Apache.
