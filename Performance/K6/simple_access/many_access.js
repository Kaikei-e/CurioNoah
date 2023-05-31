import http from 'k6/http';

export default function () {
  http.get('http://curionoah.com:4173/home');
};