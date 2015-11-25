module.exports = {
  entry: "./src/entry.js",
  output: {
    path: "./public/",
    filename: "js/bundle.js"
  },
  devtool: "source-map",
  module: {
    loaders: [
      { test: /\.scss$/, loaders: ["style", "css?sourceMap", "sass?sourceMap"] },
      { test: /\.css$/, loader: "style-loader!css-loader" },
      {
        test: /\.(jpe?g|png|gif|svg)$/i,
        loaders: [
          "file?hash=sha512&digest=hex&name=img/[hash].[ext]",
          "image-webpack?bypassOnDebug&optimizationLevel=7&interlaced=false"
        ]
      }
    ]
  }
};
