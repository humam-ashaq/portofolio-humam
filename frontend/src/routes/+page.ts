import { error } from '@sveltejs/kit';
import { PUBLIC_VITE_API_URL } from '$env/static/public';

/** @type {import('./$types').PageLoad} */
export async function load({ fetch }) {
  try {
    const res = await fetch(`${PUBLIC_VITE_API_URL}/api/projects`);

    if (!res.ok) {
      throw error(res.status, 'Could not fetch projects');
    }

    const projects = await res.json();
    return { projects };
  } catch (e) {
    console.error(e);
    // Pastikan untuk melempar error SvelteKit agar halaman error ditampilkan
    throw error(500, 'Failed to connect to the server. Is the backend running?');
  }
}