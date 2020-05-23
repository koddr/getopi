/** @jsx h */

import { h } from "preact";
import { Link } from "preact-router/match";
import style from "./style";

// Components
import Button from "../../components/form/button";
import List from "../../components/ui/list";

const Project = (props) => (
  <main class={style.project}>
    <h1>Overview</h1>
    <div class="divider-24px"></div>
    <section class={style.row_2_1}>
      <article>
        <div class={style.field}>
          <div class={style.field_name}>Research title</div>
          <div class={style.title}>Company website research</div>
        </div>
      </article>
      <aside>
        <div class={style.field_group}>
          <div class={style.field}>
            <div class={style.field_name}>Opinion reward</div>
            <div class={style.reward}>
              <span>1</span> credit
            </div>
          </div>
          <div class={style.field}>
            <Button name="Give opinion" type="primary" />
          </div>
        </div>
      </aside>
    </section>
    <div class="divider-32px"></div>
    <section class={style.row_1_2}>
      <aside>
        <div class={style.field}>
          <div class={style.field_name}>Links</div>
          <List
            type="external-link"
            items={[
              "https://github.com/koddr",
              "https://dribbble.com/koddr/test",
            ]}
          />
        </div>
      </aside>
      <article>
        <div class={style.field}>
          <div class={style.field_name}>Research description</div>
          <div>lorem dolor sit am</div>
          <div class="divider-24px"></div>
          <div class={style.field_name}>Research goals</div>
          <List
            type="check-circle"
            items={["https github.com koddr", "https dribbble.com koddr test"]}
          />
        </div>
      </article>
    </section>
    <div class="divider-48px"></div>
    <h2>Tasks to be performed</h2>
    <div class="divider-24px"></div>
  </main>
);

export default Project;
