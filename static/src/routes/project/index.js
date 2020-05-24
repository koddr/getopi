/** @jsx h */

import { h } from "preact";
import { useEffect } from "preact/hooks";
import { Link } from "preact-router/match";

// Styles
import style from "./style";

// Components
import Block from "../../components/ui/block";
import Button from "../../components/form/button";
import List from "../../components/ui/list";

const Project = (props) => {
  useEffect(() => {
    document.title = props.alias; // TODO: replace with project title
  }, [props.title]);
  return (
    <main class={style.project}>
      <h1>Overview</h1>
      <div class="divider-24px"></div>
      <section class="column_2__1_3">
        <article>
          <div class={style.item}>
            <Block
              type="title"
              label="Research title"
              content="Company website research"
            />
          </div>
        </article>
        <aside class="m__align__content_center">
          <div class={["column_1__1", style.item].join(" ")}>
            <Block
              type="reward"
              label="Opinion reward"
              content={[<span>{props.alias}</span>, " credits"]} // TODO: replace with project reward
            />
            <Button name="Give opinion" type="primary" />
          </div>
        </aside>
      </section>
      <div class="divider-32px"></div>
      <section class="column_2__0_7 align__items_start">
        <article>
          <div class={style.item}>
            <Block
              type="text"
              label="Research description"
              content="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
            />
            <Block label="Research goals">
              <List
                type="check-circle"
                items={[
                  "https github.com koddr",
                  "https dribbble.com koddr test",
                ]}
              />
            </Block>
          </div>
        </article>
        <aside>
          <div class={style.item}>
            <Block label="Links">
              <List
                type="external-link"
                items={[
                  "https://github.com/koddr",
                  "https://dribbble.com/koddr/test",
                ]}
              />
            </Block>
          </div>
          <div class={style.item}>
            <Block type="author" label="Author">
              <img src="/assets/icons/no-avatar.svg" alt="no avatar" />
              <div class="mobile__align_content_center">
                <strong>John Doe</strong>
                <div>
                  <Link href="/">@john_doe_1987</Link>
                </div>
              </div>
            </Block>
          </div>
        </aside>
      </section>
      <div class="divider-48px"></div>
      <h2>Tasks</h2>
      <div class="divider-24px"></div>
    </main>
  );
};

export default Project;
