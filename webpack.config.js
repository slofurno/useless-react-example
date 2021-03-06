module.exports = {
    entry: './src/index.js',
    output: {
        filename: './static/bundle.js'
    },
    module: {
        loaders: [
            {   test: /\.js$/,
                loader: 'babel-loader',
                query: {
                    presets: ['es2015', 'react']
                }
            }
        ]
    },
    resolve: {
        extensions: ['', '.js', '.jsx']
    }
};
