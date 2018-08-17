<template>
  <section>
    <form class="tootForm" @submit="onSubmit">
      <input v-model="instanceName" name="instance" class="instanceName" type="text" placeholder="mstdn.jp" >
      <textarea v-model="tootBody" name="body" class="preview"/>
      <input type="submit" class="submit" value="トゥート">
    </form>
    <p class="errorText">{{ error ? error : '&nbsp;' }}</p>
    <p class="poweredBy">
      Powered by <a href="/" target="_blank" class="poweredByLink">Yam</a>
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
import axios from 'axios';
import Cookies from 'js-cookie';

export default {
  data() {
    return {
      error: '',
      instanceName: '',
      tootBody: '',
    };
  },
  mounted() {
    this.tootBody = this.$route.query.text ? this.$route.query.text : '';
    this.instanceName = this.getLastInstance();
    const queryInstanceName = this.$route.query.instance;
    if (queryInstanceName) {
      this.instanceName = queryInstanceName;
      if (!this.validateForm()) {
        return;
      }
      if (
        this.$route.query.redirect === 'true' ||
        this.$route.query.redirect === 1
      ) {
        this.onCompleteInputValidation();
      }
    }
  },
  methods: {
    validateForm() {
      if (this.tootBody === '' || this.instanceName === '') {
        this.error = 'すべて入力してください。';
        return false;
      }
      return true;
    },
    onSubmit(e) {
      e.preventDefault();
      if (!this.validateForm()) {
        return;
      }
      this.onCompleteInputValidation();
    },
    onCompleteInputValidation() {
      this.prepareRedirect().then(() => {
        this.redirect();
        this.storeToCookie();
      });
    },
    prepareRedirect() {
      if (this.instanceName.startsWith('http://')) {
        this.instanceName = this.instanceName.replace('http://', '');
      }
      if (this.instanceName.startsWith('https://')) {
        this.instanceName = this.instanceName.replace('https://', '');
      }

      return new Promise((resolve, reject) => {
        const baseUrl = `https://${this.instanceName}`;
        this.pingMastodonInstance(baseUrl)
          .then(() => resolve())
          .catch(() => {
            this.error = 'エラーが発生しました。';
            reject();
          });
      });
    },
    pingMastodonInstance(url) {
      return new Promise((resolve, reject) => {
        axios
          .get(`${url}/api/v1/instance`)
          .then(() => resolve())
          .catch(() => reject());
      });
    },
    storeToCookie() {
      Cookies.set('lastInstance', this.instanceName);
    },
    getLastInstance() {
      return Cookies.get('lastInstance');
    },
    redirect() {
      if (!process.server) {
        window.location.href = `https://${
          this.instanceName
        }/share?text=${encodeURIComponent(this.tootBody)}`;
      }
    },
  },
};
</script>

<style scoped>
.tootForm {
  display: flex;
  flex-direction: column;
  margin-top: 32px;
  width: 380px;
  max-width: 85%;
}
.instanceName {
  border: none;
  margin-bottom: 12px;
}
.preview {
  resize: none;
  border: none;
  height: 96px;
}
.instanceName,
.preview {
  font-size: 1.1rem;
  border-radius: 2px;
  padding: 8px;
}

.submit {
  background: #2b90d9;
  color: #fff;
  border: none;
  margin-top: 32px;
  font-weight: bold;
  font-size: 1.2rem;
  padding: 8px 0;
  cursor: pointer;
  appearance: none;
  -webkit-appearance: none;
}

.errorText {
  font-weight: bold;
  color: red;
  margin-top: 24px;
}
</style>
