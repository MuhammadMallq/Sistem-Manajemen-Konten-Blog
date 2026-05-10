import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import MainLayout from './components/templates/MainLayout';
import Dashboard from './pages/Dashboard';
import Articles from './pages/Articles';
import Authors from './pages/Authors';
import Categories from './pages/Categories';
import Comments from './pages/Comments';
import './App.css';

function App() {
  return (
    <Router>
      <MainLayout>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/articles" element={<Articles />} />
          <Route path="/authors" element={<Authors />} />
          <Route path="/categories" element={<Categories />} />
          <Route path="/comments" element={<Comments />} />
        </Routes>
      </MainLayout>
    </Router>
  );
}

export default App;
