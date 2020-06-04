/** @jsx h */

import { h } from "preact";
import { useEffect, useState } from "preact/hooks";
import { route } from "preact-router";
import { Link } from "preact-router/match";
import { useStoreon } from "storeon/preact";
import superagent from "superagent";

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

  // Create local state
  const [project, setProject] = useState({
    status: 0,
    attrs: {},
    author: {},
  });

  useEffect(() => {
    // Load project data from API
    superagent
      .get(`http://192.168.88.100:3000/api/public/project/${props.alias}`)
      .then((res) => {
        // Append project data to local state
        setProject({
          status: res.body.project.status,
          attrs: res.body.project.project_attrs,
          author: res.body.author,
        });
      })
      .catch((err) => route(`/error/${err.status}`, true));

    // Set META attributes
    document.title = `${project.attrs.title} by @${project.author.username} → GetOpi`;
  }, [project.attrs.title, project.author.username, props.alias]);

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
          <section class="column_2__1_3 align__items_start">
            <article>
              <div class={style.item}>
                <Block
                  type="title"
                  label="Research title"
                  content={project.attrs.title}
                />
              </div>
            </article>
            <aside class="m__align__content_center">
              <div class={["column_1__1", style.item].join(" ")}>
                <Block
                  type="reward"
                  label="Opinion reward"
                  content={[
                    <span>{project.attrs.reward}</span>,
                    +project.attrs.reward === 1 ? " credit" : " credits",
                  ]}
                />
                <Button name="Give opinion" />
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
                  content={project.attrs.description}
                />
                <Block label="Research goals">
                  <List type="check-circle" items={project.attrs.goals || []} />
                </Block>
              </div>
            </article>
            <aside>
              <div class={style.item}>
                <Block label="Links">
                  <List
                    type="external-link"
                    items={project.attrs.links || []}
                  />
                </Block>
              </div>
              <div class={style.item}>
                <Block type="author" label="Author">
                  <img
                    src={
                      project.author.picture
                        ? project.author.picture
                        : "/assets/icons/no-avatar.svg"
                    }
                    alt="author avatar"
                  />
                  <div>
                    <strong>{`${project.author.first_name} ${project.author.last_name}`}</strong>
                    <div>
                      <Link href={`/user/${project.author.username}`}>
                        @{project.author.username}
                      </Link>
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
              type="success"
              default_state={showCompletedTasks}
              callback={
                () => dispatch("show completed tasks", !showCompletedTasks) // TODO: add function for show completed tasks
              }
            />
          </section>
          <div class="divider-48px" />
          <section>
            <aside>
              <div class={style.item}>
                <Block type="locked-tasks">
                  <img src="/assets/icons/lock.svg" alt="lock icon" />
                  <div>
                    Click to “Give opinion” to open{" "}
                    <strong>all research tasks</strong> and receive{" "}
                    <strong>
                      {+project.attrs.reward === 1
                        ? `${project.attrs.reward} credit`
                        : `${project.attrs.reward} credits`}
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
