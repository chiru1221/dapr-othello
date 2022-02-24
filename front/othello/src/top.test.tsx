import { render, screen } from '@testing-library/react';
import {
  BrowserRouter as Router,
}
from "react-router-dom";
import Top from './top';

test('display init setting', () => {
    render(
        <div>
            <Router >
                <Top />
            </Router>
        </div>
    );
    expect(screen.getByText(/Attaks/i)).toBeInTheDocument();
    expect(screen.getByText(/Level/i)).toBeInTheDocument();
    expect(screen.getByText(/Start/i)).toBeInTheDocument();
});
