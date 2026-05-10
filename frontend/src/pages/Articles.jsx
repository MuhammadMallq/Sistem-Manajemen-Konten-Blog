import { useState, useEffect } from 'react';
import { HiPlus, HiPencil, HiTrash } from 'react-icons/hi2';
import { getArticles, createArticle, updateArticle, deleteArticle, getAuthors, getCategories } from '../services/api';
import toast from 'react-hot-toast';

import PageHeader from '../components/templates/PageHeader';
import TableTemplate from '../components/templates/TableTemplate';
import Modal from '../components/organisms/Modal';
import ConfirmDialog from '../components/organisms/ConfirmDialog';
import SearchBar from '../components/molecules/SearchBar';
import FormGroup from '../components/molecules/FormGroup';
import Button from '../components/atoms/Button';
import Badge from '../components/atoms/Badge';
import Spinner from '../components/atoms/Spinner';

export default function Articles() {
  const [articles, setArticles] = useState([]);
  const [authors, setAuthors] = useState([]);
  const [categories, setCategories] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [showDelete, setShowDelete] = useState(null);
  const [editing, setEditing] = useState(null);
  const [search, setSearch] = useState('');
  const [form, setForm] = useState({ title: '', content: '', excerpt: '', cover_image: '', status: 'draft', author_id: '', category_id: '' });

  useEffect(() => { fetchAll(); }, []);

  const fetchAll = async () => {
    try {
      const [artRes, authRes, catRes] = await Promise.all([getArticles(), getAuthors(), getCategories()]);
      setArticles(artRes.data.data || []);
      setAuthors(authRes.data.data || []);
      setCategories(catRes.data.data || []);
    } catch (err) { toast.error('Gagal memuat data'); }
    finally { setLoading(false); }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = { ...form, author_id: Number(form.author_id), category_id: Number(form.category_id) };
    try {
      if (editing) { await updateArticle(editing.id, payload); toast.success('Artikel berhasil diperbarui'); }
      else { await createArticle(payload); toast.success('Artikel berhasil dibuat'); }
      setShowModal(false); resetForm(); fetchAll();
    } catch (err) { toast.error('Gagal menyimpan artikel'); }
  };

  const handleDelete = async () => {
    try { await deleteArticle(showDelete.id); toast.success('Artikel berhasil dihapus'); setShowDelete(null); fetchAll(); }
    catch (err) { toast.error('Gagal menghapus artikel'); }
  };

  const openEdit = (a) => {
    setEditing(a);
    setForm({ title: a.title, content: a.content, excerpt: a.excerpt || '', cover_image: a.cover_image || '', status: a.status, author_id: a.author_id, category_id: a.category_id });
    setShowModal(true);
  };

  const resetForm = () => { setForm({ title: '', content: '', excerpt: '', cover_image: '', status: 'draft', author_id: '', category_id: '' }); setEditing(null); };

  const filtered = articles.filter(a => a.title?.toLowerCase().includes(search.toLowerCase()));

  if (loading) return <Spinner />;

  return (
    <div>
      <PageHeader 
        title="Artikel" 
        description="Kelola semua artikel blog"
        actions={
          <>
            <SearchBar value={search} onChange={e => setSearch(e.target.value)} placeholder="Cari artikel..." />
            <Button variant="primary" icon={HiPlus} onClick={() => { resetForm(); setShowModal(true); }}>
              Tambah Artikel
            </Button>
          </>
        }
      />

      <TableTemplate 
        headers={['Judul', 'Penulis', 'Kategori', 'Status', 'Tanggal', 'Aksi']}
        isEmpty={filtered.length === 0}
        emptyMessage="Tidak ada artikel"
      >
        {filtered.map(a => (
          <tr key={a.id}>
            <td style={{ fontWeight: 500, maxWidth: 250, overflow: 'hidden', textOverflow: 'ellipsis', whiteSpace: 'nowrap' }}>{a.title}</td>
            <td>{a.author?.name || '-'}</td>
            <td><Badge type="category">{a.category?.name || '-'}</Badge></td>
            <td><Badge type={a.status}>{a.status}</Badge></td>
            <td style={{ fontSize: '13px', color: 'var(--text-muted)' }}>{new Date(a.created_at).toLocaleDateString('id-ID')}</td>
            <td>
              <div className="actions-cell">
                <Button variant="icon" onClick={() => openEdit(a)}><HiPencil /></Button>
                <Button variant="icon" onClick={() => setShowDelete(a)} style={{ color: 'var(--accent-red)' }}><HiTrash /></Button>
              </div>
            </td>
          </tr>
        ))}
      </TableTemplate>

      <Modal title={editing ? 'Edit Artikel' : 'Tambah Artikel Baru'} isOpen={showModal} onClose={() => setShowModal(false)}>
        <form onSubmit={handleSubmit}>
          <FormGroup label="Judul" required value={form.title} onChange={e => setForm({...form, title: e.target.value})} placeholder="Judul artikel" />
          <FormGroup label="Konten" as="textarea" required value={form.content} onChange={e => setForm({...form, content: e.target.value})} placeholder="Isi artikel..." />
          <FormGroup label="Excerpt" value={form.excerpt} onChange={e => setForm({...form, excerpt: e.target.value})} placeholder="Ringkasan singkat" />
          <FormGroup label="Cover Image URL" value={form.cover_image} onChange={e => setForm({...form, cover_image: e.target.value})} placeholder="https://..." />
          
          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr 1fr', gap: '12px' }}>
            <FormGroup label="Penulis" as="select" required value={form.author_id} onChange={e => setForm({...form, author_id: e.target.value})}>
              <option value="">Pilih Penulis</option>
              {authors.map(a => <option key={a.id} value={a.id}>{a.name}</option>)}
            </FormGroup>
            
            <FormGroup label="Kategori" as="select" required value={form.category_id} onChange={e => setForm({...form, category_id: e.target.value})}>
              <option value="">Pilih Kategori</option>
              {categories.map(c => <option key={c.id} value={c.id}>{c.name}</option>)}
            </FormGroup>
            
            <FormGroup label="Status" as="select" value={form.status} onChange={e => setForm({...form, status: e.target.value})}>
              <option value="draft">Draft</option>
              <option value="published">Published</option>
            </FormGroup>
          </div>
          
          <div className="modal-actions">
            <Button type="button" variant="secondary" onClick={() => setShowModal(false)}>Batal</Button>
            <Button type="submit" variant="primary">Simpan</Button>
          </div>
        </form>
      </Modal>

      <ConfirmDialog 
        isOpen={!!showDelete} 
        onClose={() => setShowDelete(null)}
        onConfirm={handleDelete}
        title="Hapus Artikel?"
        message={`Apakah Anda yakin ingin menghapus "<strong>${showDelete?.title}</strong>"?`}
      />
    </div>
  );
}
