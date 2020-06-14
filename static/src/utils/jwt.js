// ...
export function saveToken(token) {
  sessionStorage.setItem("tokenData", JSON.stringify(token));
}

// ...
export async function getTokenData(email, password) {
  const res = await fetch("http://192.168.88.100:5000/api/public/auth", {
    method: "POST",
    // credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email,
      password,
    }),
  });

  const tokenData = await res.json();

  if (res.status === 200) {
    // сохраняем полученный токен в sessionStorage, с помощью функции, заданной ранее
    saveToken(tokenData.auth);
    return window.location.replace(loginUrl);
  }

  return Promise.reject();
}

// ...
export function refreshToken(token) {
  return fetch("http://192.168.88.100:5000/api/private/refresh-token", {
    method: "POST",
    // credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      token,
    }),
  }).then((res) => {
    if (res.status === 200) {
      const tokenData = res.json();

      // сохраняем полученный обновленный токен в sessionStorage, с помощью функции, заданной ранее
      saveToken(JSON.stringify(tokenData));

      return Promise.resolve();
    }
    return Promise.reject();
  });
}

// ...
async function fetchWithAuth(url, options) {
  const loginUrl = "/"; // url страницы для авторизации
  let tokenData = null; // объявляем локальную переменную tokenData

  if (sessionStorage.tokenData) {
    // если в sessionStorage присутствует tokenData, то берем её
    tokenData = JSON.parse(sessionStorage.tokenData);
  } else {
    // если токен отсутствует, то перенаправляем пользователя на страницу авторизации
    return window.location.replace(loginUrl);
  }

  if (!options.headers) {
    // если в запросе отсутствует headers, то задаем их
    options.headers = {};
  }

  if (tokenData) {
    // проверяем не истек ли срок жизни токена
    if (Date.now() >= tokenData.expired_at * 1000) {
      try {
        // если истек, то обновляем токен с помощью refresh_token
        const newToken = await refreshToken(tokenData.refresh_token);
        saveToken(newToken);
      } catch {
        // если тут что-то пошло не так, то перенаправляем пользователя на страницу авторизации
        return window.location.replace(loginUrl);
      }
    }

    // добавляем токен в headers запроса
    options.headers.Authorization = `Bearer ${tokenData.access_token}`;
  }

  // возвращаем изначальную функцию, но уже с валидным токеном в headers
  return fetch(url, options);
}

export default fetchWithAuth;
