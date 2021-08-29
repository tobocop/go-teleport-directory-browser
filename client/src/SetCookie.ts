export function setCookie(cname: string, cvalue: string, ehours: number) {
  const d = new Date();
  d.setTime(d.getTime() + (ehours * 60 * 60 * 1000));
  const expires = `expires=${d.toUTCString()}`;
  document.cookie = `${cname}=${cvalue};${expires};secure`;
}
