/** @jsx h */

import { h } from "preact";
import style from "./style";

const List = (props) => {
  let listType;
  switch (props.type) {
    case "external-link":
      listType = "/assets/icons/external-link.svg";
      break;
    case "check-circle":
      listType = "/assets/icons/check-circle.svg";
      break;
  }
  return (
    <div class={style.list} role="list">
      {props.items.map((item) => {
        return (
          <div class={style.list_item} role="listitem">
            <div>
              <img src={listType} alt="list icon" />
            </div>
            <div>
              {props.type === "external-link" ? (
                <a href={item} target="_blank" rel="noreferrer">
                  {item.length > 25 ? `${item.substr(0, 23)}...` : item}
                </a>
              ) : (
                item
              )}
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default List;
