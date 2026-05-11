import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

//  DASHBOARD 
export const getDashboardStats = () => api.get('/dashboard');

//  AUTHORS (PENULIS) 
export const getAuthors = () => api.get('/authors');
export const getAuthorById = (id) => api.get(`/authors/${id}`);
export const createAuthor = (data) => api.post('/authors', data);
export const updateAuthor = (id, data) => api.put(`/authors/${id}`, data);
export const deleteAuthor = (id) => api.delete(`/authors/${id}`);

//  CATEGORIES (KATEGORI) 
export const getCategories = () => api.get('/categories');
export const getCategoryById = (id) => api.get(`/categories/${id}`);
export const createCategory = (data) => api.post('/categories', data);
export const updateCategory = (id, data) => api.put(`/categories/${id}`, data);
export const deleteCategory = (id) => api.delete(`/categories/${id}`);

//  ARTICLES (ARTIKEL) 
export const getArticles = (params) => api.get('/articles', { params });
export const getArticleById = (id) => api.get(`/articles/${id}`);
export const createArticle = (data) => api.post('/articles', data);
export const updateArticle = (id, data) => api.put(`/articles/${id}`, data);
export const deleteArticle = (id) => api.delete(`/articles/${id}`);

//  COMMENTS (KOMENTAR) 
export const getComments = (params) => api.get('/comments', { params });
export const getCommentById = (id) => api.get(`/comments/${id}`);
export const createComment = (data) => api.post('/comments', data);
export const updateComment = (id, data) => api.put(`/comments/${id}`, data);
export const deleteComment = (id) => api.delete(`/comments/${id}`);
export const getCommentsByArticle = (articleId) => api.get(`/articles/${articleId}/comments`);

export default api;
