<template>
  <section class="container">
    <h1>
      Yam
    </h1>
    <h2>Share to Mastodon, simply.</h2>
    <div
      v-clipboard:copy="url"
      v-clipboard:success="onCopy"
      v-clipboard:error="onError"
      class="howToUseBox"
    >
      <span class="url">{{ url }}</span>
      <span class="helperText">
        CLICK TO COPY
      </span>
    </div>

    <p class="notice">
      {{ notice ? notice : '&nbsp;' }}
    </p>

    <footer>
      <p class="copyright">
        <a href="https://github.com/TinyKitten/YamChromeExt" target="_blank" class="poweredByLink halfOpacity">Chrome Extension</a><br>
        Copyright &copy; 2018 TinyKitten
      </p>
    </footer>
  </section>
</template>

<script>
export default {
  data() {
    return {
      url: 'https://yam.tinykitten.me/post?text=YOUR_OWN_TEXT',
      notice: '',
    };
  },
  methods: {
    onCopy() {
      setTimeout(() => {
        this.notice = '';
      }, 3000);
      this.notice = 'URLをコピーしました。';
    },
    onError() {
      this.notice = 'あなたのブラウザはコピーに対応していません。';
    },
  },
};
</script>

<style scoped>
.howToUseBox {
  text-align: center;
  margin-top: 32px;
  background: #2b90d9;
  padding: 32px 24px;
  border-radius: 2px;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.5);
  position: relative;
  cursor: pointer;
  transition: 0.25s;
  max-width: 85%;
  overflow: hidden;
}
.url {
  font-weight: bold;
  line-height: 1.25;
  word-break: break-all;
}
.helperText {
  display: block;
  position: absolute;
  right: 0;
  bottom: 0;
  font-weight: bold;
  font-size: 0.75rem;
  padding: 4px;
  background: #282c37;
  border-radius: 4px 0 0 0;
}
.howToUseBox:hover {
  box-shadow: 0 0 24px rgba(0, 0, 0, 0.75);
}

.notice {
  margin-top: 24px;
  font-weight: bold;
}
</style>
