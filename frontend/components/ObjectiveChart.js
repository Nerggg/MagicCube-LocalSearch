// components/ObjectiveChart.js

import { Line } from 'react-chartjs-2';
import { Chart as ChartJS, LineElement, CategoryScale, LinearScale, PointElement, Tooltip, Legend } from 'chart.js';

// Register Chart.js components
ChartJS.register(LineElement, CategoryScale, LinearScale, PointElement, Tooltip, Legend);

const ObjectiveChart = ({ iterOF }) => {
    // Prepare the data for the chart
    const data = {
        labels: Array.from({ length: iterOF.length }, (_, i) => i + 1), // Generate iteration labels
        datasets: [
            {
                label: 'Objective Value',
                data: iterOF,
                borderColor: 'rgba(75, 192, 192, 1)',
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
                fill: false,
                borderWidth: 2,
                pointRadius: 3,
            }
        ]
    };

    const options = {
        responsive: true,
        plugins: {
            legend: { display: true, position: 'top' },
            tooltip: { mode: 'index', intersect: false },
        },
        scales: {
            x: { title: { display: true, text: 'Iteration' } },
            y: { title: { display: true, text: 'Objective Value' }}
        },
    };

    return <Line data={data} options={options} />;
};

export default ObjectiveChart;
