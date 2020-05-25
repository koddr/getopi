/** @jsx h */

import { h } from "preact";

const Loader = () => (
  <>
    <svg
      xmlns="http://www.w3.org/2000/svg"
      style="width: 64px; margin: 128px auto; background: none; display: block; shape-rendering: auto;"
      viewBox="0 0 100 100"
      preserveAspectRatio="xMidYMid"
    >
      <rect x="17" y="26" width="14" height="32" fill="#90B7F9">
        <animate
          attributeName="y"
          repeatCount="indefinite"
          dur="1.25s"
          calcMode="spline"
          keyTimes="0;0.5;1"
          values="6.799999999999997;26;26"
          keySplines="0 0.5 0.5 1;0 0.5 0.5 1"
          begin="-0.25s"
        />
        <animate
          attributeName="height"
          repeatCount="indefinite"
          dur="1.25s"
          calcMode="spline"
          keyTimes="0;0.5;1"
          values="86.4;48;48"
          keySplines="0 0.5 0.5 1;0 0.5 0.5 1"
          begin="-0.25s"
        />
      </rect>
      <rect x="42" y="26" width="14" height="32" fill="#90B7F9">
        <animate
          attributeName="y"
          repeatCount="indefinite"
          dur="1.25s"
          calcMode="spline"
          keyTimes="0;0.5;1"
          values="11.599999999999994;26;26"
          keySplines="0 0.5 0.5 1;0 0.5 0.5 1"
          begin="-0.125s"
        />
        <animate
          attributeName="height"
          repeatCount="indefinite"
          dur="1.25s"
          calcMode="spline"
          keyTimes="0;0.5;1"
          values="76.80000000000001;48;48"
          keySplines="0 0.5 0.5 1;0 0.5 0.5 1"
          begin="-0.125s"
        />
      </rect>
      <rect x="67" y="26" width="14" height="32" fill="#90B7F9">
        <animate
          attributeName="y"
          repeatCount="indefinite"
          dur="1.25s"
          calcMode="spline"
          keyTimes="0;0.5;1"
          values="11.599999999999994;26;26"
          keySplines="0 0.5 0.5 1;0 0.5 0.5 1"
        />
        <animate
          attributeName="height"
          repeatCount="indefinite"
          dur="1.25s"
          calcMode="spline"
          keyTimes="0;0.5;1"
          values="76.80000000000001;48;48"
          keySplines="0 0.5 0.5 1;0 0.5 0.5 1"
        />
      </rect>
    </svg>
  </>
);

export default Loader;
