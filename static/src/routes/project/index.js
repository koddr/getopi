/** @jsx h */

import { h } from "preact";
import style from "./style";

// Components
import Button from "../../components/form/button";

const Project = () => (
  <main class={style.project}>
    <h1>Overview</h1>
    <div class="divider-24px"></div>
    <article>
      <section>
        <div class={style.field}>
          <div class={style.field_name}>Research title</div>
          <div class={style.title}>Company website research</div>
        </div>
      </section>
      <section class={style.field_group}>
        <div class={style.field}>
          <div class={style.field_name}>Opinion reward</div>
          <div class={style.reward}>
            <span>100</span> credits
          </div>
        </div>
        <div class={style.field}>
          <Button name="Give opinion" type="primary" />
        </div>
      </section>
    </article>
  </main>
);

export default Project;
