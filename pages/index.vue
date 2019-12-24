<template>
  <div>
    <div class="hoge">
      <div class="chat">
        <h2 class="chatTitle">
          <span class="chatTitleStyle">Xmas</span>ぼっち掲示板
        </h2>
        <hr />
        <div style="margin-bottom: 1.1em;" v-for="(text, index) in texts" :key="index">
          <p>
            {{ index + 1 }} 名前:
            <span class="chatName">{{ text[0] }}</span>
            -
            {{ text[2] }}
          </p>
          <p class="chatText">{{ text[1] }}</p>
        </div>
        <hr />
        <div>
          <div class="chatAlertBox">
            <p>名前は20文字以下でお願いします</p>
            <p>※何も入力しない場合は匿名になります</p>
            <p>メッセージは200文字以内でお願いします</p>
          </div>
          <div class="chatFlex">
            <label>名前</label>
            <input v-model="newName" maxlength="20" style="width: 9rem;" />
          </div>
          <div class="chatFlex">
            <label>メッセージ</label>
            <textarea v-model="newComment" :rows="rows" maxlength="400" />
          </div>
          <div class="chatButtonBox">
            <button @click="add">送信</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import dayjs from "dayjs";
import "dayjs/locale/ja";
import axios from "axios";

export default {
  data() {
    return {
      texts: [],
      newName: null,
      newComment: null
    };
  },
  mounted(){
    axios
      .get(`/request`)
      .then(res => {
        console.log(res);
        console.log(res.data);
        for (let i = 0; i < res.data.length; i++) {
          let piyo = [res.data[i].name, res.data[i].comment, res.data[i].time];
          console.log(this)
          this.texts.push(piyo);
        }
        // return { texts: res.data };
      })
      .catch(e => {
        console.log(e);
      });
  },
  // localstrageを使うならこれ使う
  // mounted() {
  //   if (localStorage.getItem("texts")) {
  //     try {
  //       this.texts = JSON.parse(localStorage.getItem("texts"));
  //     } catch (e) {
  //       localStorage.removeItem("texts");
  //     }
  //   }
  // },
  methods: {
    add() {
      if (!this.newName) this.newName = "匿名希望";
      if (!this.newComment) return;

      let cratedAt = dayjs().format("YYYY-MM-D HH:mm:ss");
      let textArr = Array.of(this.newName, this.newComment, cratedAt);
      console.log(textArr);
      this.texts.push(textArr);
      console.log(this.texts)
      this.newComment = "";
      this.save(textArr);
    },
    async save(items) {
      try {
        await axios.post(`/message`, {
          name: items[0],
          comment: items[1],
          time: items[2]
        });
      } catch (err) {
        console.log(err);
      }
    }
  },
  computed: {
    rows() {
      if (this.newComment) {
        let num = this.newComment.split("\n").length;
        return num > 4 ? num : 4;
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.chat {
  min-height: 100vh;
  padding: 2rem;
  background-color: #fff;
  &Flex {
    display: flex;
    flex-direction: column;
  }
  &Title {
    font-family: fantasy;
    text-align: center;
    &Style {
      color: #dd3e54;
      background: -webkit-linear-gradient(0deg, #dd3e54, #6be585);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
    }
  }
  &Text {
    padding-left: 1rem;
    white-space: pre-wrap;
    word-wrap: break-word;
  }
  &Name {
    font-weight: bold;
  }
  &Alert {
    &Box {
      background-color: #990000;
      padding: 1rem;
      margin: 2rem 0;
      display: inline-block;
      /deep/ p {
        color: #fff;
        font-weight: bold;
      }
    }
  }
  &Button {
    &Box {
      display: flex;
      justify-content: flex-end;
    }
  }
}

hr {
  border: none;
  border-top: dashed 1px #ff0000;
  height: 1px;
  color: #ffffff;
  margin: 1rem 0;
}

button {
  display: inline-block;
  margin: 1.5rem 1.5rem 0 0;
  border-radius: 5%; /* 角丸       */
  font-size: 18pt; /* 文字サイズ */
  text-align: center; /* 文字位置   */
  cursor: pointer; /* カーソル   */
  padding: 12px 12px; /* 余白       */
  background: #990000; /* 背景色     */
  color: #ffffff; /* 文字色     */
  line-height: 1em; /* 1行の高さ  */
  transition: 0.3s; /* なめらか変化 */
  box-shadow: 6px 6px 3px #666666; /* 影の設定 */
  border: 2px solid #990000; /* 枠の指定 */
}

button:hover {
  box-shadow: none; /* カーソル時の影消去 */
  color: #990000; /* 背景色     */
  background: #ffffff; /* 文字色     */
}
</style>
