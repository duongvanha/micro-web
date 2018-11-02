module.exports = {
    devServer   : {
        disableHostCheck: true,
    },
    chainWebpack: config => {
        // GraphQL Loader
        config.module
            .rule('graphql')
            .test(/\.graphql$/)
            .use('graphql-tag/loader')
            .loader('graphql-tag/loader')
            .end();
    },
};
