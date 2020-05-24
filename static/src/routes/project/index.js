/** @jsx h */

import { h } from "preact";
import { Link } from "preact-router/match";
import style from "./style";

// Components
import Block from "../../components/ui/block";
import Button from "../../components/form/button";
import List from "../../components/ui/list";

const Project = (props) => (
  <main class={style.project}>
    <h1>Overview</h1>
    <div class="divider-24px"></div>
    <section class={style.row_2_1}>
      <article>
        <Block
          type="title"
          label="Research title"
          content="Company website research"
        />
      </article>
      <aside>
        <div class={style.field_group}>
          <Block
            type="reward"
            label="Opinion reward"
            content={[<span>{props.alias}</span>, " credits"]}
          />
          <Button name="Give opinion" type="primary" />
        </div>
      </aside>
    </section>
    <div class="divider-32px"></div>
    <section class={style.row_1_2}>
      <aside>
        <Block label="Links">
          <List
            type="external-link"
            items={[
              "https://github.com/koddr",
              "https://dribbble.com/koddr/test",
            ]}
          />
        </Block>
      </aside>
      <article>
        <Block
          type="text"
          label="Research description"
          content="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
        />
        <div class="divider-32px"></div>
        <Block label="Research goals">
          <List
            type="check-circle"
            items={["https github.com koddr", "https dribbble.com koddr test"]}
          />
        </Block>
      </article>
    </section>
    <div class="divider-48px"></div>
    <h2>Tasks</h2>
    <div class="divider-24px"></div>
  </main>
);

export default Project;
