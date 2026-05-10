import { useState, useEffect } from 'react';
import { HiPlus, HiPencil, HiTrash } from 'react-icons/hi2';
import { getCategories, createCategory, updateCategory, deleteCategory } from '../services/api';
import toast from 'react-hot-toast';

import PageHeader from '../components/templates/PageHeader';
import TableTemplate from '../components/templates/TableTemplate';
import Modal from '../components/organisms/Modal';
import ConfirmDialog from '../components/organisms/ConfirmDialog';
import FormGroup from '../components/molecules/FormGroup';
import Button from '../components/atoms/Button';
import Spinner from '../components/atoms/Spinner';

export default function Categories() {
  const [categories, setCategories] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [showDelete, setShowDelete] = useState(null);
  const [editing, setEditing] = useState(null);
  const [form, setForm] = useState({ name: '', description: '', color: '#6366F1' });

  useEffect(() => { fetchData(); }, []);

  const fetchData = async () => {
    try { const res = await getCategories(); setCategories(res.data.data || []); }
    catch { toast.error('Gagal memuat kategori'); }
    finally { setLoading(false); }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (editing) { await updateCategory(editing.id, form); toast.success('Kategori berhasil diperbarui'); }
      else { await createCategory(form); toast.success('Kategori berhasil ditambahkan'); }
      setShowModal(false); resetForm(); fetchData();
    } catch { toast.error('Gagal menyimpan kategori'); }
  };

  const handleDelete = async () => {
    try { await deleteCategory(showDelete.id); toast.success('Kategori berhasil dihapus'); setShowDelete(null); fetchData(); }
    catch { toast.error('Gagal menghapus kategori'); }
  };

  const openEdit = (c) => { setEditing(c); setForm({ name: c.name, description: c.description || '', color: c.color || '#6366F1' }); setShowModal(true); };
  const resetForm = () => { setForm({ name: '', description: '', color: '#6366F1' }); setEditing(null); };

  if (loading) return <Spinner />;

  return (
    <div>
      <PageHeader 
        title="Kategori" 
        description="Kelola kategori artikel blog"
        actions={
          <Button variant="primary" icon={HiPlus} onClick={() => { resetForm(); setShowModal(true); }}>
            Tambah Kategori
          </Button>
        }
      />

      <TableTemplate 
        headers={['Warna', 'Nama', 'Deskripsi', 'Tanggal Dibuat', 'Aksi']}
        isEmpty={categories.length === 0}
        emptyMessage="Tidak ada kategori"
      >
        {categories.map(c => (
          <tr key={c.id}>
            <td><div style={{ width: 24, height: 24, borderRadius: 6, background: c.color || '#6366F1' }}></div></td>
            <td style={{ fontWeight: 500 }}>{c.name}</td>
            <td style={{ color: 'var(--text-muted)', maxWidth: 300, overflow: 'hidden', textOverflow: 'ellipsis', whiteSpace: 'nowrap' }}>{c.description || '-'}</td>
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

      <Modal title={editing ? 'Edit Kategori' : 'Tambah Kategori Baru'} isOpen={showModal} onClose={() => setShowModal(false)}>
        <form onSubmit={handleSubmit}>
          <FormGroup label="Nama Kategori" required value={form.name} onChange={e => setForm({...form, name: e.target.value})} placeholder="Nama kategori" />
          <FormGroup label="Deskripsi" as="textarea" value={form.description} onChange={e => setForm({...form, description: e.target.value})} placeholder="Deskripsi kategori..." />
          <FormGroup label="Warna" type="color" value={form.color} onChange={e => setForm({...form, color: e.target.value})} style={{ height: 44, cursor: 'pointer' }} />
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
        title="Hapus Kategori?"
        message={`Apakah Anda yakin ingin menghapus "<strong>${showDelete?.name}</strong>"?`}
      />
    </div>
  );
}
