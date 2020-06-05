/** @jsx h */

import { h } from "preact";
import style from "./style";

const Block = (props) => {
  let blockType, blockContent;
  switch (props.type) {
    case "title":
      blockType = style.block__title;
      blockContent = props.content;
      break;
    case "text":
      blockType = style.block__text;
      blockContent = props.content;
      break;
    case "reward":
      blockType = style.block__reward;
      blockContent = props.content;
      break;
    case "author":
      blockType = style.block__author;
      blockContent = props.children; // throw all children elements into <Block/>
      break;
    case "locked-tasks":
      blockType = style.block__locked_tasks;
      blockContent = props.children; // throw all children elements into <Block/>
      break;
    default:
      blockContent = props.children; // throw all children elements into <Block/>
      break;
  }
  return (
    <div class={style.block}>
      <div class={style.block__label}>{props.label}</div>
      <div class={blockType}>{blockContent}</div>
    </div>
  );
};

export default Block;
