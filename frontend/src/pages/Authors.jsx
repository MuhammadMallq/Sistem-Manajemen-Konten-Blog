import { useState, useEffect } from 'react';
import { HiPlus, HiPencil, HiTrash } from 'react-icons/hi2';
import { getAuthors, createAuthor, updateAuthor, deleteAuthor } from '../services/api';
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

export default function Authors() {
  const [authors, setAuthors] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [showDelete, setShowDelete] = useState(null);
  const [editing, setEditing] = useState(null);
  const [search, setSearch] = useState('');
  const [form, setForm] = useState({ name: '', email: '', bio: '', avatar: '' });

  useEffect(() => { fetchData(); }, []);

  const fetchData = async () => {
    try { const res = await getAuthors(); setAuthors(res.data.data || []); }
    catch { toast.error('Gagal memuat data penulis'); }
    finally { setLoading(false); }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (editing) { await updateAuthor(editing.id, form); toast.success('Penulis berhasil diperbarui'); }
      else { await createAuthor(form); toast.success('Penulis berhasil ditambahkan'); }
      setShowModal(false); resetForm(); fetchData();
    } catch { toast.error('Gagal menyimpan data'); }
  };

  const handleDelete = async () => {
    try { await deleteAuthor(showDelete.id); toast.success('Penulis berhasil dihapus'); setShowDelete(null); fetchData(); }
    catch { toast.error('Gagal menghapus penulis'); }
  };

  const openEdit = (a) => { setEditing(a); setForm({ name: a.name, email: a.email, bio: a.bio || '', avatar: a.avatar || '' }); setShowModal(true); };
  const resetForm = () => { setForm({ name: '', email: '', bio: '', avatar: '' }); setEditing(null); };
  const filtered = authors.filter(a => a.name?.toLowerCase().includes(search.toLowerCase()));

  if (loading) return <Spinner />;

  return (
    <div>
      <PageHeader 
        title="Penulis" 
        description="Kelola data penulis blog"
        actions={
          <>
            <SearchBar value={search} onChange={e => setSearch(e.target.value)} placeholder="Cari penulis..." />
            <Button variant="primary" icon={HiPlus} onClick={() => { resetForm(); setShowModal(true); }}>
              Tambah Penulis
            </Button>
          </>
        }
      />

      <TableTemplate 
        headers={['Nama', 'Email', 'Bio', 'Jumlah Artikel', 'Aksi']}
        isEmpty={filtered.length === 0}
        emptyMessage="Tidak ada penulis"
      >
        {filtered.map(a => (
          <tr key={a.id}>
            <td>
              <div style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
                <img src={a.avatar || `https://ui-avatars.com/api/?name=${a.name}&background=6366f1&color=fff`} alt="" style={{ width: 36, height: 36, borderRadius: '50%' }} />
                <span style={{ fontWeight: 500 }}>{a.name}</span>
              </div>
            </td>
            <td style={{ color: 'var(--text-secondary)' }}>{a.email}</td>
            <td style={{ maxWidth: 200, overflow: 'hidden', textOverflow: 'ellipsis', whiteSpace: 'nowrap', color: 'var(--text-muted)' }}>{a.bio || '-'}</td>
            <td><Badge type="category">{a.articles?.length || 0} artikel</Badge></td>
            <td>
              <div className="actions-cell">
                <Button variant="icon" onClick={() => openEdit(a)}><HiPencil /></Button>
                <Button variant="icon" onClick={() => setShowDelete(a)} style={{ color: 'var(--accent-red)' }}><HiTrash /></Button>
              </div>
            </td>
          </tr>
        ))}
      </TableTemplate>

      <Modal title={editing ? 'Edit Penulis' : 'Tambah Penulis Baru'} isOpen={showModal} onClose={() => setShowModal(false)}>
        <form onSubmit={handleSubmit}>
          <FormGroup label="Nama" required value={form.name} onChange={e => setForm({...form, name: e.target.value})} placeholder="Nama penulis" />
          <FormGroup label="Email" type="email" required value={form.email} onChange={e => setForm({...form, email: e.target.value})} placeholder="email@contoh.com" />
          <FormGroup label="Bio" as="textarea" value={form.bio} onChange={e => setForm({...form, bio: e.target.value})} placeholder="Biografi singkat..." />
          <FormGroup label="Avatar URL" value={form.avatar} onChange={e => setForm({...form, avatar: e.target.value})} placeholder="https://..." />
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
        title="Hapus Penulis?"
        message={`Apakah Anda yakin ingin menghapus "<strong>${showDelete?.name}</strong>"?`}
      />
    </div>
  );
}
