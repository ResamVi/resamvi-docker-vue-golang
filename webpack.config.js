const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const UglifyJsPlugin = require('uglifyjs-webpack-plugin')

module.exports = {
    entry: './src/index.js',
    
    output: {
        path: path.resolve(__dirname, 'build'),
        filename: 'index.js'
    },
    
    devtool: 'source-map',
    
    module: {
        rules: [
            {
                test: /\.vue/,
                use: ['vue-loader']
            },
            {
                test: /\.css/,
                use: ExtractTextPlugin.extract({
                    fallback: 'style-loader',
                    use: {
                        loader: 'css-loader',
                        options: {
                            minimize: true
                        }
                    }
                })
            },
            {
                test: /\.(gif|png|jpe?g|ico)/,
                use: [{
                    loader: 'file-loader',
                    options: {
                        name: '[name].[ext]',
                        outputPath: 'img/'
                    }
                }]
            },
            {
                test: /\.(woff2|woff|otf|ttf)/,
                use: [{
                    loader: 'file-loader',
                    options: {
                        name: '[name].[ext]',
                        outputPath: 'fonts/'
                    }                    
                }]
            }
        ]
    },
    plugins: [
        new ExtractTextPlugin('style.css'),
        
        new HtmlWebpackPlugin({
            filename: 'index.html',
            template: './src/index.html',
            inject: 'body',
            minify: {
                html5: false,
                removeComments: true, 
                collapseWhitespace: true,
            }
        }),
        
        new UglifyJsPlugin()
    ],

    resolve: {
        alias: {
            vue: 'vue/dist/vue.js'
        }
    }
};