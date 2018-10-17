'use strict';

module.exports = appInfo => {
  const config = exports = {};

  // use for cookie sign key, should change to your own and keep security
  config.keys = appInfo.name + '_1539745858143_8562';

  // add your config here
  config.middleware = [];

  // 添加 view 配置
  config.view = {
      defaultViewEngine: 'nunjucks',
      mapping: {
        '.tpl': 'nunjucks',
      },
    };
  
    //add news config
    config.news = {
      pageSize: 1,
      serverUrl: 'http://api.douban.com/v2/movie/top250',
    };

  return config;
};
