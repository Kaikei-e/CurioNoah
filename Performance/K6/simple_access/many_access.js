import http from 'k6/http';

export function many_access() {
  http.get('http://curionoah.com:4173/home');
};