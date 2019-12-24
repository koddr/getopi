/**
 * Validation
 * @param {String} input_type
 * @param {String} what_match
 */
function Validation(input_type, what_match) {
  let re;

  switch (input_type) {
    case "email":
      re = /^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$/;
      break;
    default:
      break;
  }

  return re.test(what_match);
}

export default Validation;
