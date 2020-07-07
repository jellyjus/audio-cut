module.exports = {
    publicPath: process.env.NODE_ENV === 'production'
        ? '/static'
        : '/',
    configureWebpack: {
        devServer: {
            watchOptions: {
                ignored: /node_modules/
            }
        }
    }

};