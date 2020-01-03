const path = require('path')

module.exports = {
    entry: './static/tour/tour.js',
    output: {
        filename: 'tour-bundle.js',
        path: path.resolve(__dirname, 'assets/assets/js')
    },
};