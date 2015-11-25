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
      { test: /\.(png|jpg)$/, loader: 'url-loader?limit=8192' }
    ]
  }
};
