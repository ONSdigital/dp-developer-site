require("@babel/polyfill");
const path = require('path')

module.exports = {
    entry: ['@babel/polyfill', './js/main.js'],
    output: {
        filename: 'bundle.js',
        path: path.resolve(__dirname, 'assets/assets/js')
    },
    module: {
        rules: [
          {
            test: /\.m?js$/,
            exclude: /(node_modules|bower_components)/,
            use: {
              loader: 'babel-loader',
              options: {
                presets: ['@babel/preset-env']
              }
            }
          }
        ]
      }
};