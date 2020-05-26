/** @jsx h */

import { h } from "preact";
import { useEffect } from "preact/hooks";
import { Link } from "preact-router/match";
import { useStoreon } from "storeon/preact";

// Styles
import style from "./style";

// Components
import Header from "../../components/header";
import Footer from "../../components/footer";
import Block from "../../components/ui/block";
import Button from "../../components/form/button";
import Checkbox from "../../components/form/checkbox";
import List from "../../components/ui/list";

const Project = (props) => {
  // Connect to store
  const { dispatch, showCompletedTasks } = useStoreon("showCompletedTasks");

  // Set META attributes
  useEffect(() => {
    document.title = props.alias; // TODO: replace with project title
  }, [props.alias, props.title]);

  // Render component
  return (
    <>
      <Header />
      <main class={style.project}>
        <div class={style.wrapper}>
          <div class={style.heading}>
            <img
              onClick={() => history.back()} // back to prev page from browser history
              src="/assets/icons/back.svg"
              alt="back icon"
            />
            <h1>Overview</h1>
          </div>
          <div class="divider-48px" />
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
                  content={[
                    <span>{props.alias}</span>,
                    +props.alias === 1 ? " credit" : " credits", // TODO: replace with project reward
                  ]}
                />
                <Button name="Give opinion" type="primary" />
              </div>
            </aside>
          </section>
          <div class="divider-32px" />
          <section class="column_2__0_7 align__items_start">
            <article>
              <div class={style.item}>
                <Block
                  type="text"
                  label="Research description"
                  content="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
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
                  <img src="/assets/icons/no-avatar.svg" alt="no avatar icon" />
                  <div>
                    <strong>John Doe</strong>
                    <div>
                      <Link href="/">@john_doe_1987</Link>
                    </div>
                  </div>
                </Block>
              </div>
            </aside>
          </section>
          <div class="divider-48px" />
          <section class={style.tasks}>
            <h2>Tasks</h2>
            <Checkbox
              name="check"
              text="Show completed tasks"
              default_state={showCompletedTasks}
              callback={() => {
                dispatch("show completed tasks", !showCompletedTasks);
                console.log("TODO: add function for show completed tasks");
              }}
            />
          </section>
          <div class="divider-24px" />
          <section>
            <aside>
              <div class={style.item}>
                <Block type="locked-tasks">
                  <img src="/assets/icons/lock.svg" alt="lock icon" />
                  <div>
                    Click to “Give opinion” to open{" "}
                    <strong>all research tasks</strong> and receive{" "}
                    <strong>
                      {+props.alias === 1
                        ? `${props.alias} credit`
                        : `${props.alias} credits`}
                    </strong>
                    !
                  </div>
                  <Button name="Give opinion" type="secondary" />
                </Block>
              </div>
            </aside>
          </section>
        </div>
      </main>
      <Footer />
    </>
  );
};

export default Project;
