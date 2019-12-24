<template>
  <div>
    <div class="container mx-auto" id="app">
      <h2 class="text-3xl">ひと（り）ごと掲示板</h2>
      <hr />
      <div style="margin-bottom: 1.1em;" v-for="(text,n) in texts" :key="text">
        <p>
          <span class="text">名前: {{text[0]}}「{{text[1]}} - {{text[2]}}」</span>
          <button
            class="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow"
            @click="remove(n)"
          >Remove</button>
        </p>
      </div>

      <hr />
      <div>
        <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">名前：</label>
        <input
          class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
          v-model="newName"
        />
        <br />
        <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">ひとこと：</label>
        <input
          class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
          v-model="newComment"
        />
        <br />
        <button
          style="margin-top: 1.5em;"
          class="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow"
          @click="add"
        >送信</button>
      </div>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs';
import 'dayjs/locale/ja';

export default {
  data() {
    return {
      texts: [],
      newName: null,
      newComment: null
    };
  },
  mounted() {
    if (localStorage.getItem("texts")) {
      try {
        this.texts = JSON.parse(localStorage.getItem("texts"));
      } catch (e) {
        localStorage.removeItem("texts");
      }
    }
  },
  methods: {
    add() {
      if (!this.newName) this.newName = "匿名希望";
      if (!this.newComment) return;

      let cratedAt = dayjs().format("YYYY-MM-D HH:mm:ss");
      let textArr = Array.of(this.newName, this.newComment, cratedAt);
      this.texts.push(textArr);
      this.newName = "";
      this.newComment = "";
      this.save();
    },
    remove(x) {
      this.texts.splice(x, 1);
      this.save();
    },
    save() {
      let parsed = JSON.stringify(this.texts);
      localStorage.setItem("texts", parsed);
    }
  }
};
</script>