const path = require("path");

module.exports = {
  mode: "development", // Set to 'production' for production build
  entry: "./src/index.ts", // Entry point of your application

  output: {
    path: path.resolve(__dirname, "dist"), // Output directory
    filename: "bundle.js", // Output bundle file name
  },

  resolve: {
    extensions: [".ts", ".tsx", ".js"], // Resolve TypeScript and JavaScript files
  },

  module: {
    rules: [
      {
        test: /\.tsx?$/, // Use ts-loader for TypeScript files
        exclude: /node_modules/,
        use: "ts-loader",
      },
    ],
  },
};
