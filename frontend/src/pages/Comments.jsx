import { useState, useEffect } from 'react';
import { HiPlus, HiPencil, HiTrash } from 'react-icons/hi2';
import { getComments, createComment, updateComment, deleteComment, getArticles } from '../services/api';
import toast from 'react-hot-toast';

import PageHeader from '../components/templates/PageHeader';
import TableTemplate from '../components/templates/TableTemplate';
import Modal from '../components/organisms/Modal';
import ConfirmDialog from '../components/organisms/ConfirmDialog';
import FormGroup from '../components/molecules/FormGroup';
import Button from '../components/atoms/Button';
import Badge from '../components/atoms/Badge';
import Spinner from '../components/atoms/Spinner';

export default function Comments() {
  const [comments, setComments] = useState([]);
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [showDelete, setShowDelete] = useState(null);
  const [editing, setEditing] = useState(null);
  const [form, setForm] = useState({ article_id: '', commenter_name: '', commenter_email: '', content: '' });

  useEffect(() => { fetchAll(); }, []);

  const fetchAll = async () => {
    try {
      const [comRes, artRes] = await Promise.all([getComments(), getArticles()]);
      setComments(comRes.data.data || []);
      setArticles(artRes.data.data || []);
    } catch { toast.error('Gagal memuat data'); }
    finally { setLoading(false); }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = { ...form, article_id: Number(form.article_id) };
    try {
      if (editing) { await updateComment(editing.id, payload); toast.success('Komentar berhasil diperbarui'); }
      else { await createComment(payload); toast.success('Komentar berhasil ditambahkan'); }
      setShowModal(false); resetForm(); fetchAll();
    } catch { toast.error('Gagal menyimpan komentar'); }
  };

  const handleDelete = async () => {
    try { await deleteComment(showDelete.id); toast.success('Komentar berhasil dihapus'); setShowDelete(null); fetchAll(); }
    catch { toast.error('Gagal menghapus komentar'); }
  };

  const openEdit = (c) => { setEditing(c); setForm({ article_id: c.article_id, commenter_name: c.commenter_name, commenter_email: c.commenter_email, content: c.content }); setShowModal(true); };
  const resetForm = () => { setForm({ article_id: '', commenter_name: '', commenter_email: '', content: '' }); setEditing(null); };

  if (loading) return <Spinner />;

  return (
    <div>
      <PageHeader 
        title="Komentar" 
        description="Kelola komentar pada artikel blog"
        actions={
          <Button variant="primary" icon={HiPlus} onClick={() => { resetForm(); setShowModal(true); }}>
            Tambah Komentar
          </Button>
        }
      />

      <TableTemplate 
        headers={['Nama', 'Email', 'Komentar', 'Artikel', 'Tanggal', 'Aksi']}
        isEmpty={comments.length === 0}
        emptyMessage="Tidak ada komentar"
      >
        {comments.map(c => (
          <tr key={c.id}>
            <td style={{ fontWeight: 500 }}>{c.commenter_name}</td>
            <td style={{ color: 'var(--text-secondary)' }}>{c.commenter_email}</td>
            <td style={{ maxWidth: 250, overflow: 'hidden', textOverflow: 'ellipsis', whiteSpace: 'nowrap', color: 'var(--text-muted)' }}>{c.content}</td>
            <td><Badge type="category">{c.article?.title?.substring(0, 30) || '-'}</Badge></td>
            <td style={{ fontSize: '13px', color: 'var(--text-muted)' }}>{new Date(c.created_at).toLocaleDateString('id-ID')}</td>
            <td>
              <div className="actions-cell">
                <Button variant="icon" onClick={() => openEdit(c)}><HiPencil /></Button>
                <Button variant="icon" onClick={() => setShowDelete(c)} style={{ color: 'var(--accent-red)' }}><HiTrash /></Button>
              </div>
            </td>
          </tr>
        ))}
      </TableTemplate>

      <Modal title={editing ? 'Edit Komentar' : 'Tambah Komentar Baru'} isOpen={showModal} onClose={() => setShowModal(false)}>
        <form onSubmit={handleSubmit}>
          <FormGroup label="Artikel" as="select" required value={form.article_id} onChange={e => setForm({...form, article_id: e.target.value})}>
            <option value="">Pilih Artikel</option>
            {articles.map(a => <option key={a.id} value={a.id}>{a.title}</option>)}
          </FormGroup>
          <FormGroup label="Nama" required value={form.commenter_name} onChange={e => setForm({...form, commenter_name: e.target.value})} placeholder="Nama komentator" />
          <FormGroup label="Email" type="email" required value={form.commenter_email} onChange={e => setForm({...form, commenter_email: e.target.value})} placeholder="email@contoh.com" />
          <FormGroup label="Komentar" as="textarea" required value={form.content} onChange={e => setForm({...form, content: e.target.value})} placeholder="Tulis komentar..." />
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
        title="Hapus Komentar?"
        message={`Apakah Anda yakin ingin menghapus komentar dari "<strong>${showDelete?.commenter_name}</strong>"?`}
      />
    </div>
  );
}
