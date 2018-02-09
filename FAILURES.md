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
* [kahun/awesome-sysadmin](https://github.com/kahun/awesome-sysadmin/blob/master/LICENSE.txt) - CC-BY-SA 4.0, clear failure.
* [BVLC/caffe](https://github.com/BVLC/caffe/blob/master/LICENSE) - custom header and footer.
* [JuliaLang/julia](https://github.com/JuliaLang/julia/blob/master/LICENSE.md) - custom header and huge footer with dependencies description.
* [AFNetworking/AFNetworking](https://github.com/AFNetworking/AFNetworking/blob/master/LICENSE) - MIT, clear failure.
* [eslint/eslint](https://github.com/eslint/eslint/blob/master/LICENSE) - MIT, clear failure.
* [browserify/browserify](https://github.com/browserify/browserify/blob/master/LICENSE) - several licenses concatenated.
* [Theano/Theano](https://github.com/Theano/Theano/blob/master/LICENSE.txt) - symlink to `doc/LICENSE.txt`.
* [facebookarchive/three20](https://github.com/facebookarchive/three20/blob/master/LICENSE) - points to Apache.
* [torch/torch7](https://github.com/torch/torch7/blob/master/COPYRIGHT.txt) - `COPYRIGHT.txt` with custom header.
* [tmux/tmux](https://github.com/tmux/tmux) - custom `COPYING`, BSD is mentioned in `README`.
* [aFarkas/html5shiv](https://github.com/aFarkas/html5shiv/blob/master/MIT%20and%20GPL2%20licenses.md) - the file is `MIT and GPL2 licenses.md` and it is a concatenation.
* [atlassian/localstack](https://github.com/atlassian/localstack/blob/master/LICENSE.txt) - points to Apache.
* [dmlc/xgboost](https://github.com/dmlc/xgboost/blob/master/LICENSE) - points to Apache.
* [facebook/nuclide](https://github.com/facebook/nuclide/blob/master/LICENSE) - custom license? resembles BSD.
* [mathiasbynens/dotfiles](https://github.com/mathiasbynens/dotfiles/blob/master/LICENSE-MIT.txt) - the file is `LICENSE-MIT.txt`.
* [realm/realm-java](https://github.com/realm/realm-java/blob/master/LICENSE) - custom format, Apache concatenated with various notices.
* [google/deepdream](https://github.com/google/deepdream/blob/master/LICENSE) - Apache with custom header.
* [Microsoft/CNTK](https://github.com/Microsoft/CNTK/blob/master/LICENSE.md) - custom format, concatenation.
* [robbiehanson/CocoaAsyncSocket](https://github.com/robbiehanson/CocoaAsyncSocket/blob/master/LICENSE.txt) - public domain and BSD, custom format.
* [vinta/awesome-python](https://github.com/vinta/awesome-python/blob/master/LICENSE) - points to CC-BY 4.0.
* [pytorch/pytorch](https://github.com/pytorch/pytorch/blob/master/LICENSE) - custom header before BSD.
* [EbookFoundation/free-programming-books](https://github.com/EbookFoundation/free-programming-books/blob/master/LICENSE) - points to CC-BY 4.0.
* [stedolan/jq](https://github.com/stedolan/jq/blob/master/COPYING) - concatenation of several licenses.
* [phanan/htaccess](https://github.com/phanan/htaccess/blob/master/LICENSE) - custom, says something about public domain and unlicense.
* [strongloop/loopback](https://github.com/strongloop/loopback/blob/master/LICENSE) - MIT with custom header.
* [MaximAbramchuck/awesome-interview-questions](https://github.com/MaximAbramchuck/awesome-interview-questions) - README states CC0.
* [enaqx/awesome-react](https://github.com/enaqx/awesome-react#license) - README states CC0.
* [CocoaLumberjack/CocoaLumberjack](https://github.com/CocoaLumberjack/CocoaLumberjack/blob/master/LICENSE.txt) - BSD, clear failure.
* [Modernizr/Modernizr](https://github.com/Modernizr/Modernizr) - license file is a joke, README mentions MIT.
* [aosabook/500lines](https://github.com/aosabook/500lines/blob/master/LICENSE.md) - custom format, BSD in the end.
* [videojs/video.js](https://github.com/videojs/video.js/blob/master/LICENSE) - points to Apache.
* [jquery/jquery-ui](https://github.com/jquery/jquery-ui/blob/master/LICENSE.txt) - concatenation of JQuery and CC0.
* [Automattic/mongoose](https://github.com/Automattic/mongoose) - BSD license is appended to the end of `README`.
* [alibaba/druid](https://github.com/alibaba/druid/blob/master/license.txt) - points to Apache.
* [vuejs/awesome-vue](https://github.com/vuejs/awesome-vue) - CC0 is appended to the end of `README`.
* [mxgmn/WaveFunctionCollapse](https://github.com/mxgmn/WaveFunctionCollapse/blob/master/LICENSE.md) - custom text, points to MIT.
* [janpaepke/ScrollMagic](https://github.com/janpaepke/ScrollMagic/blob/master/LICENSE.md) - MIT with custom header, mentions GPL.
* [marionettejs/backbone.marionette](https://github.com/marionettejs/backbone.marionette/blob/master/license.txt) - custom text, mentions MIT, links to http://mutedsolutions.mit-license.org/.
* [Marak/faker.js](https://github.com/Marak/faker.js/blob/master/MIT-LICENSE.txt) - file is `MIT-LICENSE.txt`, custom header.
* [astaxie/beego](https://github.com/astaxie/beego/blob/master/LICENSE) - points to Apache.
* [requests/requests](https://github.com/requests/requests/blob/master/LICENSE) - points to Apache.
* [donnemartin/interactive-coding-challenges](https://github.com/donnemartin/interactive-coding-challenges/blob/master/LICENSE) - points to Apache.
* [cockroachdb/cockroach](https://github.com/cockroachdb/cockroach/blob/master/LICENSE) - custom text, mentions Apache, CCL, MIT and BSD, 
* [linnovate/mean](https://github.com/linnovate/mean/blob/master/LICENSE) - MIT with broken lines, clear failure.
* [sqlmapproject/sqlmap](https://github.com/sqlmapproject/sqlmap/blob/master/LICENSE) - GPL with custom formatting, not difficult.
* [bcit-ci/CodeIgniter](https://github.com/bcit-ci/CodeIgniter/blob/develop/license.txt) - MIT, clear failure.
