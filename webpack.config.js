module.exports = {
  entry: "src/entry.js",
  output: {
    path: __dirname,
    filename: "public/js/bundle.js"
  },
  module: {
    loaders: [
      { test: /\.css$/, loader: "style!css" },
      { test: /\.scss$/, loaders: ["style", "css", "sass"] }
    ]
  }
};
