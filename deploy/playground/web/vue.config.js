module.exports = {
  transpileDependencies: ["vuetify"],
  devServer: {
    proxy: {
      "/api": {
        target: "http://localhost:5000/",
        changeOrigin: true
      },
      "/halborn": {
        target: "http://localhost:1317/",
        changeOrigin: true
      }
    }
  }
};
