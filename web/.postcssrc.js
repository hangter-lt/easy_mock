/**
 * PostCSS 配置文件
 */

module.exports = {
    // 配置要使用的 PostCSS 插件
    plugins: {
        // 配置使用 postcss-pxtorem 插件
        // 作用：把 px 转为 rem
        'postcss-pxtorem': {
            mediaQuery: false, //媒体查询( @media screen 之类的)中不生效
            minPixelValue: 5, //px小于12的不会被转换
            rootValue: 192, // 或者直接指定所有的都为37.5 二选一
            propList: ['*']
        }
    }
}
