import { useState, useEffect } from 'react';
import { HiOutlineDocumentText, HiOutlineUsers, HiOutlineTag, HiOutlineChatBubbleLeftRight } from 'react-icons/hi2';
import { getDashboardStats } from '../services/api';
import PageHeader from '../components/templates/PageHeader';
import StatCard from '../components/molecules/StatCard';
import Spinner from '../components/atoms/Spinner';
import Badge from '../components/atoms/Badge';

export default function Dashboard() {
  const [stats, setStats] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchStats();
  }, []);

  const fetchStats = async () => {
    try {
      const res = await getDashboardStats();
      setStats(res.data.data);
    } catch (err) {
      console.error('Error fetching stats:', err);
    } finally {
      setLoading(false);
    }
  };

  if (loading) return <Spinner />;

  return (
    <div>
      <PageHeader 
        title="Dashboard" 
        description="Selamat datang di Sistem Manajemen Konten Blog" 
      />

      <div className="stats-grid">
        <StatCard title="Total Artikel" value={stats?.total_articles || 0} icon={HiOutlineDocumentText} />
        <StatCard title="Total Penulis" value={stats?.total_authors || 0} icon={HiOutlineUsers} />
        <StatCard title="Total Kategori" value={stats?.total_categories || 0} icon={HiOutlineTag} />
        <StatCard title="Total Komentar" value={stats?.total_comments || 0} icon={HiOutlineChatBubbleLeftRight} />
      </div>

      <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '24px' }}>
        <div className="card">
          <h3 style={{ fontSize: '16px', fontWeight: 600, marginBottom: '16px' }}>Artikel Terbaru</h3>
          {stats?.recent_articles?.length > 0 ? stats.recent_articles.map(a => (
            <div key={a.id} style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '12px 0', borderBottom: '1px solid var(--border-color)' }}>
              <div>
                <div style={{ fontWeight: 500, fontSize: '14px' }}>{a.title}</div>
                <div style={{ fontSize: '12px', color: 'var(--text-muted)', marginTop: 2 }}>{a.author?.name} • {a.category?.name}</div>
              </div>
              <Badge type={a.status}>{a.status}</Badge>
            </div>
          )) : <p className="empty-state" style={{padding: '20px'}}>Belum ada artikel</p>}
        </div>
        <div className="card">
          <h3 style={{ fontSize: '16px', fontWeight: 600, marginBottom: '16px' }}>Komentar Terbaru</h3>
          {stats?.recent_comments?.length > 0 ? stats.recent_comments.map(c => (
            <div key={c.id} style={{ padding: '12px 0', borderBottom: '1px solid var(--border-color)' }}>
              <div style={{ fontWeight: 500, fontSize: '14px' }}>{c.commenter_name}</div>
              <div style={{ fontSize: '13px', color: 'var(--text-secondary)', marginTop: 4 }}>{c.content?.substring(0, 80)}...</div>
            </div>
          )) : <p className="empty-state" style={{padding: '20px'}}>Belum ada komentar</p>}
        </div>
      </div>
    </div>
  );
}
