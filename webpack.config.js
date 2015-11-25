module.exports = {
  entry: "./src/entry.js",
  output: {
    path: __dirname,
    filename: "public/js/bundle.js"
  },
  devtool: "source-map",
  module: {
    loaders: [
      { test: /\.scss$/, loaders: ["style", "css?sourceMap", "sass?sourceMap"] },
      { test: /\.css$/, loader: 'style-loader!css-loader' },
      {
        test: /\.(jpe?g|png|gif|svg)$/i,
        loaders: [
          'file?hash=sha512&digest=hex&name=public/img/[hash].[ext]',
          'image-webpack?bypassOnDebug&optimizationLevel=7&interlaced=false'
        ]
      }
    ]
  }
};
