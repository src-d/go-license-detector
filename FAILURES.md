# License detection failures

This is a list of known license detection failures. Once there appear common patterns, go-license-detector
will be extended to support those.

### Clear failures
* [Microsoft/vscode](https://github.com/Microsoft/vscode/blob/master/LICENSE.txt) - license file seems legit, clear failure.
* [adobe/brackets](https://github.com/adobe/brackets/blob/master/LICENSE) - license file seems legit, clear failure.
* [pure-css/pure](https://github.com/pure-css/pure/blob/master/LICENSE.md) - BSD license not detected, clear failure.
* [TryGhost/Ghost](https://github.com/TryGhost/Ghost/blob/master/LICENSE) - MIT license, clear failure.
* [nwjs/nw.js](https://github.com/nwjs/nw.js/blob/nw28/LICENSE) - MIT, clear failure.
* [jenkinsci/jenkins](https://github.com/jenkinsci/jenkins/blob/master/LICENSE.txt) - MIT license, clear failure.
* [kahun/awesome-sysadmin](https://github.com/kahun/awesome-sysadmin/blob/master/LICENSE.txt) - CC-BY-SA 4.0, clear failure.
* [AFNetworking/AFNetworking](https://github.com/AFNetworking/AFNetworking/blob/master/LICENSE) - MIT, clear failure.
* [eslint/eslint](https://github.com/eslint/eslint/blob/master/LICENSE) - MIT, clear failure.
* [CocoaLumberjack/CocoaLumberjack](https://github.com/CocoaLumberjack/CocoaLumberjack/blob/master/LICENSE.txt) - BSD, clear failure.
* [linnovate/mean](https://github.com/linnovate/mean/blob/master/LICENSE) - MIT with broken lines, clear failure.
* [bcit-ci/CodeIgniter](https://github.com/bcit-ci/CodeIgniter/blob/develop/license.txt) - MIT, clear failure.
* [moment/moment](https://github.com/moment/moment/blob/develop/LICENSE) - MIT, clear failure.
* [webpack/webpack](https://github.com/webpack/webpack/blob/master/LICENSE) - MIT, clear failure.
* [Alamofire/Alamofire](https://github.com/Alamofire/Alamofire/blob/master/LICENSE) - MIT, clear failure.

### Custom file names
* [atech/postal](https://github.com/atech/postal) - the file is `MIT-LICENCE`.
* [rust-lang/rust](https://github.com/rust-lang/rust) - license files: `LICENSE-APACHE` and `LICENSE-MIT`; also the `README` tells about them.
* [ariya/phantomjs](https://github.com/ariya/phantomjs/blob/master/LICENSE.BSD) - the file is `LICENSE.BSD`.
* [mpv-player/mpv](https://github.com/mpv-player/mpv) - `LICENSE.GPL` and `LICENSE.LGPL`.
* [v8/v8](https://github.com/v8/v8) - `LICENSE.v8` and 4 other custom suffixes; `LICENSE` is a concatenation.
* [philc/vimium](https://github.com/philc/vimium/blob/master/MIT-LICENSE.txt) - `MIT-LICENSE.txt`.
* [torch/torch7](https://github.com/torch/torch7/blob/master/COPYRIGHT.txt) - `COPYRIGHT.txt` with custom header.
* [mathiasbynens/dotfiles](https://github.com/mathiasbynens/dotfiles/blob/master/LICENSE-MIT.txt) - the file is `LICENSE-MIT.txt`.
* [Marak/faker.js](https://github.com/Marak/faker.js/blob/master/MIT-LICENSE.txt) - file is `MIT-LICENSE.txt`, custom header.
* [gionkunz/chartist-js](https://github.com/gionkunz/chartist-js/blob/develop/LICENSE-MIT) - the file is `LICENSE-MIT`.
* [php-fig/fig-standards](https://github.com/php-fig/fig-standards) - the file names are `LICENSE-CC.md` and `LICENSE-MIT.md`.
* [VundleVim/Vundle.vim](https://github.com/VundleVim/Vundle.vim/blob/master/LICENSE-MIT.txt) - the file name is `LICENSE-MIT.txt`.
* [carhartl/jquery-cookie](https://github.com/carhartl/jquery-cookie/blob/master/MIT-LICENSE.txt) - the file name is `MIT-LICENSE.txt`.
* [JetBrains/kotlin](https://github.com/JetBrains/kotlin/tree/master/license) - `license` directory, the standard file there points to Apache.

### Pointers
* [akullpp/awesome-java](https://github.com/akullpp/awesome-java/blob/master/LICENSE.md) - file consists of the single reference to CC-BY-SA-4.0.
* [Unitech/pm2](https://github.com/Unitech/pm2/blob/master/LICENSE) - the whole content is the name of the real license file.
* [ruanyf/es6tutorial](https://github.com/ruanyf/es6tutorial/blob/gh-pages/LICENSE) - human-readable summary of CC-BY-NC 4.0.
* [lukasz-madon/awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job) - `README` contains a link to CC0.
* [sindresorhus/quick-look-plugins](https://github.com/sindresorhus/quick-look-plugins) - `README` points to CC0.
* [MaximAbramchuck/awesome-interview-questions](https://github.com/MaximAbramchuck/awesome-interview-questions) - `README` states CC0.
* [enaqx/awesome-react](https://github.com/enaqx/awesome-react#license) - `README` states CC0.
* [sorrycc/awesome-javascript](https://github.com/sorrycc/awesome-javascript#license) - `README` points to CC0.
* [vuejs/awesome-vue](https://github.com/vuejs/awesome-vue) - CC0 is appended to the end of `README`.
* [terryum/awesome-deep-learning-papers](https://github.com/terryum/awesome-deep-learning-papers) - `README` points to CC0.
* [gztchan/awesome-design](https://github.com/gztchan/awesome-design) - `README` points to CC0.
* [sindresorhus/awesome-electron](https://github.com/sindresorhus/awesome-electron) - `README` points to CC0.
* [donnemartin/system-design-primer](https://github.com/donnemartin/system-design-primer/blob/master/LICENSE.txt) - custom text, link to CC-BY 4.0.
* [vinta/awesome-python](https://github.com/vinta/awesome-python/blob/master/LICENSE) - points to CC-BY 4.0.
* [EbookFoundation/free-programming-books](https://github.com/EbookFoundation/free-programming-books/blob/master/LICENSE) - points to CC-BY 4.0.
* [kamranahmedse/design-patterns-for-humans](https://github.com/kamranahmedse/design-patterns-for-humans#license) - `README` points to CC-BY 4.0.
* [interagent/http-api-design](https://github.com/interagent/http-api-design/blob/master/LICENSE.md) - license file points to CC-BY 3.0.
* [kamranahmedse/developer-roadmap](https://github.com/kamranahmedse/developer-roadmap) - `README` points to CC-BY 3.0.
* [saltstack/salt](https://github.com/saltstack/salt/blob/develop/LICENSE) - custom text, link to Apache.
* [inconshreveable/ngrok](https://github.com/inconshreveable/ngrok/blob/master/LICENSE) - points to Apache.
* [google/python-fire](https://github.com/google/python-fire/blob/master/LICENSE) - points to Apache.
* [swagger-api/swagger-ui](https://github.com/swagger-api/swagger-ui/blob/master/LICENSE) - points to Apache.
* [SignalR/SignalR](https://github.com/SignalR/SignalR/blob/dev/LICENSE.txt) - points to Apache.
* [facebookarchive/three20](https://github.com/facebookarchive/three20/blob/master/LICENSE) - points to Apache.
* [atlassian/localstack](https://github.com/atlassian/localstack/blob/master/LICENSE.txt) - points to Apache.
* [dmlc/xgboost](https://github.com/dmlc/xgboost/blob/master/LICENSE) - points to Apache.
* [videojs/video.js](https://github.com/videojs/video.js/blob/master/LICENSE) - points to Apache.
* [alibaba/druid](https://github.com/alibaba/druid/blob/master/license.txt) - points to Apache.
* [astaxie/beego](https://github.com/astaxie/beego/blob/master/LICENSE) - points to Apache.
* [requests/requests](https://github.com/requests/requests/blob/master/LICENSE) - points to Apache.
* [alibaba/fastjson](https://github.com/alibaba/fastjson/blob/master/license.txt) - file points to Apache.
* [Reactive-Extensions/RxJS](https://github.com/Reactive-Extensions/RxJS/blob/master/license.txt) - points to Apache.
* [lord/slate](https://github.com/lord/slate/blob/master/LICENSE) - points to Apache.
* [donnemartin/data-science-ipython-notebooks](https://github.com/donnemartin/data-science-ipython-notebooks/blob/master/LICENSE) - points to Apache.
* [donnemartin/interactive-coding-challenges](https://github.com/donnemartin/interactive-coding-challenges/blob/master/LICENSE) - points to Apache.
* [SFTtech/openage](https://github.com/SFTtech/openage) - `README` mentions GNU GPLv3.
* [FFmpeg/FFmpeg](https://github.com/FFmpeg/FFmpeg/blob/master/LICENSE.md) - custom license text, points to LGPL and GPL.
* [androidannotations/androidannotations](https://github.com/androidannotations/androidannotations/blob/develop/LICENSE.txt) - license points to Apache and CDDL.
* [Theano/Theano](https://github.com/Theano/Theano/blob/master/LICENSE.txt) - symlink to `doc/LICENSE.txt`.
* [mxgmn/WaveFunctionCollapse](https://github.com/mxgmn/WaveFunctionCollapse/blob/master/LICENSE.md) - custom text, points to MIT.
* [marionettejs/backbone.marionette](https://github.com/marionettejs/backbone.marionette/blob/master/license.txt) - custom text, mentions MIT, links to http://mutedsolutions.mit-license.org/.
* [date-fns/date-fns](https://github.com/date-fns/date-fns/blob/master/LICENSE.md) - custom text, points to http://kossnocorp.mit-license.org/.
* [shadowsocks/shadowsocks-android](https://github.com/shadowsocks/shadowsocks-android/blob/master/LICENSE) - shortened GPL.
* [mozilla/BrowserQuest](https://github.com/mozilla/BrowserQuest/blob/master/LICENSE) - file points to MPL and CC-BY-SA 3.0.

### Headers and footers
* [Carthage/Carthage](https://github.com/Carthage/Carthage/blob/master/LICENSE.md) - extra content with copyright at the bottom of MIT license.
* [serverless/serverless](https://github.com/serverless/serverless/blob/master/LICENSE.txt) - MIT with extra notice at the bottom.
* [gitlabhq/gitlabhq](https://github.com/gitlabhq/gitlabhq/blob/master/LICENSE) - extra content at the end of MIT or BSD license.
* [akveo/blur-admin](https://github.com/akveo/blur-admin/blob/master/LICENSE.txt) - MIT license with garbage in the end.
* [keras-team/keras](https://github.com/keras-team/keras/blob/master/LICENSE) - MIT with many copyright notices in the beginning.
* [meteor/meteor](https://github.com/meteor/meteor/blob/devel/LICENSE) - custom notice appended in the end of MIT license.
* [celery/celery](https://github.com/celery/celery/blob/master/LICENSE) - BSD with custom header and footer.
* [parse-community/parse-server](https://github.com/parse-community/parse-server/blob/master/LICENSE) - custom footer after BSD.
* [ccgus/fmdb](https://github.com/ccgus/fmdb/blob/master/LICENSE.txt) - custom header before MIT.
* [BVLC/caffe](https://github.com/BVLC/caffe/blob/master/LICENSE) - custom header and footer.
* [JuliaLang/julia](https://github.com/JuliaLang/julia/blob/master/LICENSE.md) - custom header and huge footer with dependencies description.
* [google/deepdream](https://github.com/google/deepdream/blob/master/LICENSE) - Apache with custom header.
* [pytorch/pytorch](https://github.com/pytorch/pytorch/blob/master/LICENSE) - custom header before BSD.
* [strongloop/loopback](https://github.com/strongloop/loopback/blob/master/LICENSE) - MIT with custom header.
* [aosabook/500lines](https://github.com/aosabook/500lines/blob/master/LICENSE.md) - custom format, BSD in the end.
* [janpaepke/ScrollMagic](https://github.com/janpaepke/ScrollMagic/blob/master/LICENSE.md) - MIT with custom header, mentions GPL.
* [google/protobuf](https://github.com/google/protobuf/blob/master/LICENSE) - BSD with custom header.
* [rapid7/metasploit-framework](https://github.com/rapid7/metasploit-framework/blob/master/COPYING) - BSD with custom footer.
* [ipython/ipython](https://github.com/ipython/ipython/blob/master/COPYING.rst) - custom format, huge footer.
* [jquery/jquery](https://github.com/jquery/jquery/blob/master/LICENSE.txt) - JQuery with custom footer (he-he).
* [hapijs/hapi](https://github.com/hapijs/hapi/blob/master/LICENSE) - BSD with custom header and especially footer.
* [sqlmapproject/sqlmap](https://github.com/sqlmapproject/sqlmap/blob/master/LICENSE) - GPL with custom header.

### Concatenations
* [nodejs/node](https://github.com/nodejs/node/blob/master/LICENSE) - multiple licenses in the same file.
* [chrissimpkins/Hack](https://github.com/source-foundry/Hack/blob/master/LICENSE.md) - multiple licenses in the same file.
* [shadowsocks/shadowsocks-windows](https://github.com/shadowsocks/shadowsocks-windows/blob/master/LICENSE.txt) - multiple licenses concatenated together.
* [Tencent/mars](https://github.com/Tencent/mars/blob/master/LICENSE) - many licenses concatenated, custom header.
* [lodash/lodash](https://github.com/lodash/lodash/blob/master/LICENSE) - several licenses concatenated.
* [libuv/libuv](https://github.com/libuv/libuv/blob/v1.x/LICENSE) - several licenses concatenated together.
* [iview/iview](https://github.com/iview/iview/blob/2.0/LICENSE) - several MIT licenses concatenated together.
* [fatih/vim-go](https://github.com/fatih/vim-go/blob/master/LICENSE) - two BSDs concatenated together, custom header, middle and bottom. 
* [bumptech/glide](https://github.com/bumptech/glide/blob/master/LICENSE) - concatenated licenses.
* [sqlitebrowser/sqlitebrowser](https://github.com/sqlitebrowser/sqlitebrowser/blob/master/LICENSE) - GPL and MPL concatenated.
* [Mantle/Mantle](https://github.com/Mantle/Mantle/blob/master/LICENSE.md) - two licenses concatenated.
* [kripken/emscripten](https://github.com/kripken/emscripten/blob/incoming/LICENSE) - several licenses concatenated, custom header and footer.
* [browserify/browserify](https://github.com/browserify/browserify/blob/master/LICENSE) - several licenses concatenated.
* [aFarkas/html5shiv](https://github.com/aFarkas/html5shiv/blob/master/MIT%20and%20GPL2%20licenses.md) - the file is `MIT and GPL2 licenses.md` and it is a concatenation.
* [Microsoft/CNTK](https://github.com/Microsoft/CNTK/blob/master/LICENSE.md) - custom format, concatenation.
* [stedolan/jq](https://github.com/stedolan/jq/blob/master/COPYING) - concatenation of several licenses.
* [jquery/jquery-ui](https://github.com/jquery/jquery-ui/blob/master/LICENSE.txt) - concatenation of JQuery and CC0.
* [realm/realm-cocoa](https://github.com/realm/realm-cocoa/blob/master/LICENSE) - Apache concatenated with other content.
* [robbiehanson/CocoaAsyncSocket](https://github.com/robbiehanson/CocoaAsyncSocket/blob/master/LICENSE.txt) - public domain and BSD, custom format.
* [jquery/jquery-mobile](https://github.com/jquery/jquery-mobile/blob/master/LICENSE.txt) - concatenation of JQuery and CC0.

### Die hards
* [RubaXa/Sortable](https://github.com/RubaXa/Sortable) - license is appended to the end of `README.md`.
* [Automattic/mongoose](https://github.com/Automattic/mongoose) - BSD license is appended to the end of `README`.
* [mperham/sidekiq](https://github.com/mperham/sidekiq/blob/master/LICENSE) - license file is completely custom, mentions LGPL.
* [opencv/opencv](https://github.com/opencv/opencv/blob/master/LICENSE) - license is completely custom format but resembles a BSD.
* [python/cpython](https://github.com/python/cpython/blob/master/LICENSE) - license is PSF 2.
* [tmux/tmux](https://github.com/tmux/tmux) - custom `COPYING`, BSD is mentioned in `README`.
* [facebook/nuclide](https://github.com/facebook/nuclide/blob/master/LICENSE) - custom license? resembles BSD.
* [realm/realm-java](https://github.com/realm/realm-java/blob/master/LICENSE) - custom format, Apache concatenated with various notices.
* [phanan/htaccess](https://github.com/phanan/htaccess/blob/master/LICENSE) - custom, says something about public domain and unlicense.
* [cockroachdb/cockroach](https://github.com/cockroachdb/cockroach/blob/master/LICENSE) - custom text, mentions Apache, CCL, MIT and BSD, 
* [fbsamples/f8app](https://github.com/fbsamples/f8app/blob/master/LICENSE) - customized MIT.
* [FortAwesome/Font-Awesome](https://github.com/FortAwesome/Font-Awesome/blob/master/LICENSE.txt) - custom format, mentions several licenses.
* [Microsoft/api-guidelines](https://github.com/Microsoft/api-guidelines/blob/master/Guidelines.md#44-license) - pointer to CC-BY 4.0 in `Guidelines.md`.
* [neovim/neovim](https://github.com/neovim/neovim/blob/master/LICENSE) - Apache with custom header and notes.
* [moklick/frontend-stuff](https://github.com/moklick/frontend-stuff/blob/master/LICENSE) - summarized CC0.

### Miscellaneous
* [IanLunn/Hover](https://github.com/IanLunn/Hover/blob/master/license.txt) - custom license, nothing can be done.
* [xdissent/ievms](https://github.com/xdissent/ievms) - "none" license stated in the `README`.
* [isocpp/CppCoreGuidelines](https://github.com/isocpp/CppCoreGuidelines/blob/master/LICENSE) - custom license.
* [Modernizr/Modernizr](https://github.com/Modernizr/Modernizr) - license file is a joke, `README` mentions MIT.
* [froala/design-blocks](https://github.com/froala/design-blocks/blob/dev/LICENSE) - FROALA OPEN WEB DESIGN LICENSE.
* [Swordfish90/cool-retro-term](https://github.com/Swordfish90/cool-retro-term) - `gpl-3.0.txt` and `gpl-2.0.txt` files.
