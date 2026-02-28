// API 基础路径
const API_BASE = '/api';

// 端口数据缓存
let portsData = [];

// 当前过滤状态
let currentFilter = 'all';

// 当前语言
let currentLang = 'en';

// 应用模板配置
const appTemplates = [
    { name: 'MySQL', port: 3306, description: 'MySQL Database', icon: '🗄️' },
    { name: 'Redis', port: 6379, description: 'Redis Cache', icon: '🔴' },
    { name: 'PostgreSQL', port: 5432, description: 'PostgreSQL Database', icon: '🐘' },
    { name: 'MongoDB', port: 27017, description: 'MongoDB Database', icon: '🍃' },
    { name: 'Nginx', port: 80, description: 'Nginx Web Server', icon: '🌐' },
    { name: 'Apache', port: 80, description: 'Apache Web Server', icon: '🪶' },
    { name: 'Tomcat', port: 8080, description: 'Tomcat Application Server', icon: '🐱' },
    { name: 'Node.js', port: 3000, description: 'Node.js Application', icon: '💚' },
    { name: 'Vue', port: 8080, description: 'Vue Development Server', icon: '💚' },
    { name: 'React', port: 3000, description: 'React Development Server', icon: '⚛️' },
    { name: 'Jenkins', port: 8080, description: 'Jenkins CI/CD', icon: '👷' },
    { name: 'GitLab', port: 80, description: 'GitLab Server', icon: '🦊' },
    { name: 'RabbitMQ', port: 5672, description: 'RabbitMQ Message Queue', icon: '🐰' },
    { name: 'Kafka', port: 9092, description: 'Kafka Message Queue', icon: '📨' },
    { name: 'Elasticsearch', port: 9200, description: 'Elasticsearch Search Engine', icon: '🔍' },
    { name: 'Docker', port: 2375, description: 'Docker API', icon: '🐳' },
    { name: 'SSH', port: 22, description: 'SSH Server', icon: '🔐' },
    { name: 'FTP', port: 21, description: 'FTP Server', icon: '📁' },
    { name: 'DNS', port: 53, description: 'DNS Server', icon: '🌐' },
    { name: 'SMTP', port: 25, description: 'SMTP Mail Server', icon: '📧' },
];

// 语言包
const i18n = {
    en: {
        title: 'Port Radar',
        searchPlaceholder: 'Search ports or apps...',
        scan: 'Scan',
        occupiedPorts: 'Occupied Ports',
        marked: 'Marked',
        port: 'Port',
        protocol: 'Protocol',
        process: 'Process',
        localAddr: 'Local Address',
        appMark: 'App Mark',
        actions: 'Actions',
        loading: 'Loading...',
        scanning: 'Scanning ports...',
        noData: 'No data',
        loadFailed: 'Load failed: ',
        editMark: 'Edit Mark',
        addMark: 'Add Mark',
        appName: 'App Name',
        appNamePlaceholder: 'e.g., Nginx, MySQL',
        description: 'Description',
        descriptionPlaceholder: 'App description or notes...',
        cancel: 'Cancel',
        save: 'Save',
        markSaved: 'Mark saved',
        saveFailed: 'Save failed',
        markDeleted: 'Mark deleted',
        deleteFailed: 'Delete failed',
        confirmDelete: 'Are you sure to delete this mark?',
        unmarked: 'Unmarked',
        edit: 'Edit',
        mark: 'Mark',
        removeMark: 'Unmark',
        appTemplate: 'App Template (click to apply)',
        kill: 'Kill',
        confirmKill: 'Are you sure to kill this process? This may cause data loss.',
        processKilled: 'Process killed successfully',
        killFailed: 'Failed to kill process',
        processNotFound: 'Process not found after kill, refreshing...',
        templateApplied: 'Template applied',
        // Docker 相关
        runningContainers: 'Running Containers',
        container: 'Container',
        containerDetail: 'Container Detail',
        containerName: 'Container Name',
        containerImage: 'Image',
        containerState: 'State',
        containerPorts: 'Port Mappings',
        stopContainer: 'Stop',
        startContainer: 'Start',
        restartContainer: 'Restart',
        removeContainer: 'Remove',
        confirmStop: 'Are you sure to stop this container?',
        confirmStart: 'Are you sure to start this container?',
        confirmRestart: 'Are you sure to restart this container?',
        confirmRemove: 'Are you sure to remove this container? This action cannot be undone!',
        containerStopped: 'Container stopped',
        containerStarted: 'Container started',
        containerRestarted: 'Container restarted',
        containerRemoved: 'Container removed',
        containerActionFailed: 'Container action failed',
        close: 'Close',
        dockerPort: 'Docker Port'
    },
    zh: {
        title: '端口雷达',
        searchPlaceholder: '搜索端口或应用...',
        scan: '扫描',
        occupiedPorts: '已占用端口',
        marked: '已标记',
        port: '端口',
        protocol: '协议',
        process: '进程',
        localAddr: '本地地址',
        appMark: '应用标记',
        actions: '操作',
        loading: '加载中...',
        scanning: '扫描端口中...',
        noData: '暂无数据',
        loadFailed: '加载失败: ',
        editMark: '编辑标记',
        addMark: '添加标记',
        appName: '应用名称',
        appNamePlaceholder: '例如: Nginx, MySQL',
        description: '描述',
        descriptionPlaceholder: '应用描述或备注...',
        cancel: '取消',
        save: '保存',
        markSaved: '标记已保存',
        saveFailed: '保存失败',
        markDeleted: '标记已删除',
        deleteFailed: '删除失败',
        confirmDelete: '确定要删除此标记吗？',
        unmarked: '未标记',
        edit: '编辑',
        mark: '标记',
        removeMark: '取消标记',
        appTemplate: '应用模板（点击应用）',
        kill: '终止',
        confirmKill: '确定要终止此进程吗？可能导致数据丢失！',
        processKilled: '进程已终止',
        killFailed: '终止进程失败',
        processNotFound: '进程已终止，正在刷新...',
        templateApplied: '模板已应用',
        // Docker 相关
        runningContainers: '运行中容器',
        container: '容器',
        containerDetail: '容器详情',
        containerName: '容器名称',
        containerImage: '镜像',
        containerState: '状态',
        containerPorts: '端口映射',
        stopContainer: '停止',
        startContainer: '启动',
        restartContainer: '重启',
        removeContainer: '删除',
        confirmStop: '确定要停止此容器吗？',
        confirmStart: '确定要启动此容器吗？',
        confirmRestart: '确定要重启此容器吗？',
        confirmRemove: '确定要删除此容器吗？此操作不可恢复！',
        containerStopped: '容器已停止',
        containerStarted: '容器已启动',
        containerRestarted: '容器已重启',
        containerRemoved: '容器已删除',
        containerActionFailed: '容器操作失败',
        close: '关闭',
        dockerPort: 'Docker端口'
    }
};

// DOM 元素
const portListEl = document.getElementById('portList');
const searchInput = document.getElementById('searchInput');
const refreshBtn = document.getElementById('refreshBtn');
const modal = document.getElementById('modal');
const markForm = document.getElementById('markForm');
const containerModal = document.getElementById('containerModal');

// Docker 状态缓存
let dockerStats = { available: false, totalContainers: 0, runningContainers: 0 };

// 初始化
document.addEventListener('DOMContentLoaded', () => {
    // 初始化语言
    initLanguage();

    // 渲染模板选择器
    renderTemplates();

    loadPorts();
    loadDockerStats();

    refreshBtn.addEventListener('click', () => {
        loadPorts();
        loadDockerStats();
    });
    searchInput.addEventListener('input', handleSearch);
    markForm.addEventListener('submit', handleSaveMark);
});

// ==================== 国际化相关 ====================

// 初始化语言
function initLanguage() {
    // 尝试从 localStorage 获取
    const savedLang = localStorage.getItem('portManagerLang');

    if (savedLang && (savedLang === 'en' || savedLang === 'zh')) {
        currentLang = savedLang;
    } else {
        // 根据浏览器语言自动检测
        const browserLang = navigator.language || navigator.userLanguage;
        if (browserLang && browserLang.toLowerCase().startsWith('zh')) {
            currentLang = 'zh';
        } else {
            currentLang = 'en';
        }
    }

    applyLanguage();
}

// 切换语言
function toggleLanguage() {
    currentLang = currentLang === 'en' ? 'zh' : 'en';
    localStorage.setItem('portManagerLang', currentLang);
    applyLanguage();
    applyFilter(); // 重新渲染端口列表
}

// 应用语言
function applyLanguage() {
    const langBtn = document.getElementById('langToggle');
    langBtn.textContent = currentLang === 'en' ? '🌐 EN' : '🌐 中文';

    // 更新页面标题
    document.title = t('title');

    // 更新所有 data-i18n 元素
    document.querySelectorAll('[data-i18n]').forEach(el => {
        const key = el.dataset.i18n;
        if (i18n[currentLang][key]) {
            el.textContent = i18n[currentLang][key];
        }
    });

    // 更新所有 data-i18n-placeholder 元素
    document.querySelectorAll('[data-i18n-placeholder]').forEach(el => {
        const key = el.dataset.i18nPlaceholder;
        if (i18n[currentLang][key]) {
            el.placeholder = i18n[currentLang][key];
        }
    });

    // 更新 html lang 属性
    document.documentElement.lang = currentLang === 'zh' ? 'zh-CN' : 'en';
}

// 获取翻译文本
function t(key) {
    return i18n[currentLang][key] || key;
}

// ==================== 应用模板相关 ====================

// 渲染模板选择器
function renderTemplates() {
    const grid = document.getElementById('templateGrid');
    grid.innerHTML = appTemplates.map(tpl => `
        <div class="template-item" onclick="applyTemplate('${tpl.name}', '${tpl.description}')" title="${tpl.name} (${tpl.port})">
            <span class="template-icon">${tpl.icon}</span>
            <span class="template-name">${tpl.name}</span>
        </div>
    `).join('');
}

// 应用模板
function applyTemplate(name, description) {
    document.getElementById('appName').value = name;
    document.getElementById('description').value = description;
    showToast(t('templateApplied'));
}

// ==================== 端口列表相关 ====================

// 加载端口列表
async function loadPorts() {
    portListEl.innerHTML = `<tr><td colspan="7" class="loading">
        <div class="scanning-indicator">
            <div class="scanning-spinner"></div>
            <span>${t('scanning')}</span>
        </div>
    </td></tr>`;

    try {
        const res = await fetch(`${API_BASE}/ports`);
        portsData = await res.json();

        // 按端口号升序排序
        portsData.sort((a, b) => a.port - b.port);

        // 保持当前过滤状态
        if (currentFilter !== 'all') {
            document.querySelector(`.stat-item[data-filter="${currentFilter}"]`).classList.add('active');
        }

        applyFilter();
        updateStats(portsData);
    } catch (err) {
        portListEl.innerHTML = `<tr><td colspan="7" class="loading" style="color: var(--danger);">${t('loadFailed')}${err.message}</td></tr>`;
    }
}

// 渲染端口列表
function renderPorts(ports) {
    if (ports.length === 0) {
        portListEl.innerHTML = `<tr><td colspan="7" class="loading">${t('noData')}</td></tr>`;
        return;
    }

    portListEl.innerHTML = ports.map(port => {
        // 构建进程名显示
        let processDisplay = escapeHtml(port.processName || '-');
        if (port.container) {
            processDisplay = `<span class="process-name docker-process" title="${t('dockerPort')}: ${port.container.name}">
                🐳 ${escapeHtml(port.container.name)}
            </span>`;
        }

        // 构建应用标记显示
        let appMarkDisplay = '';
        if (port.container) {
            appMarkDisplay = `<span class="app-mark docker-mark" title="${port.container.image}">
                🐳 ${escapeHtml(port.container.name)}
            </span>`;
        } else if (port.appMark) {
            appMarkDisplay = `<span class="app-mark" title="${escapeHtml(port.appMark.description || '')}">${escapeHtml(port.appMark.appName)}</span>`;
        } else {
            appMarkDisplay = `<span class="app-mark empty">${t('unmarked')}</span>`;
        }

        // 构建操作按钮
        let actions = `
            <button class="btn btn-small btn-primary" onclick="openEditModal(${port.port}, '${port.protocol}', ${port.appMark ? 'true' : 'false'})">
                ${port.appMark ? t('edit') : t('mark')}
            </button>
        `;

        if (port.appMark && !port.container) {
            actions += `<button class="btn btn-small btn-secondary" onclick="deleteMark(${port.port}, '${port.protocol}')">${t('removeMark')}</button>`;
        }

        // Docker容器操作按钮
        if (port.container) {
            const containerId = port.container.id;
            const containerName = escapeHtml(port.container.name);
            const isRunning = port.container.state === 'running';
            
            if (isRunning) {
                actions += `
                    <button class="btn btn-small btn-warning" onclick="dockerAction('${containerId}', 'stop', '${containerName}')">${t('stopContainer')}</button>
                    <button class="btn btn-small btn-secondary" onclick="dockerAction('${containerId}', 'restart', '${containerName}')">${t('restartContainer')}</button>
                `;
            } else {
                actions += `
                    <button class="btn btn-small btn-success" onclick="dockerAction('${containerId}', 'start', '${containerName}')">${t('startContainer')}</button>
                    <button class="btn btn-small btn-danger" onclick="dockerAction('${containerId}', 'remove', '${containerName}')">${t('removeContainer')}</button>
                `;
            }
        } else if (port.pid && port.pid > 0) {
            actions += `<button class="btn btn-small btn-danger" onclick="killProcess(${port.pid}, '${escapeHtml(port.processName || '')}')">${t('kill')}</button>`;
        }

        return `
            <tr class="${port.container ? 'docker-row' : ''}">
                <td><span class="port-number">${port.port}</span></td>
                <td><span class="protocol-tag ${port.protocol}">${port.protocol}</span></td>
                <td>${processDisplay}</td>
                <td><span class="pid">${port.container ? '🐳' : (port.pid || '-')}</span></td>
                <td><span class="pid">${escapeHtml(port.localAddr || '-')}</span></td>
                <td>${appMarkDisplay}</td>
                <td class="actions">${actions}</td>
            </tr>
        `;
    }).join('');
}

// 更新统计数据
function updateStats(ports) {
    const marked = ports.filter(p => p.appMark).length;
    const docker = ports.filter(p => p.container).length;
    const tcp = ports.filter(p => p.protocol === 'tcp').length;
    const udp = ports.filter(p => p.protocol === 'udp').length;

    document.getElementById('totalPorts').textContent = ports.length;
    document.getElementById('markedPorts').textContent = marked;
    document.getElementById('dockerPorts').textContent = docker;
    document.getElementById('tcpPorts').textContent = tcp;
    document.getElementById('udpPorts').textContent = udp;
}

// 加载Docker统计
async function loadDockerStats() {
    try {
        const res = await fetch(`${API_BASE}/docker/stats`);
        dockerStats = await res.json();

        // 显示Docker容器统计
        if (dockerStats.available) {
            document.getElementById('dockerStatItem').style.display = 'block';
            document.getElementById('dockerContainers').textContent = dockerStats.runningContainers;
        } else {
            document.getElementById('dockerStatItem').style.display = 'none';
        }
    } catch (err) {
        console.log('Docker stats not available:', err);
        document.getElementById('dockerStatItem').style.display = 'none';
    }
}

// ==================== 过滤相关 ====================

// 过滤处理
function handleFilter(filter) {
    currentFilter = filter;

    // 更新激活状态
    document.querySelectorAll('.stat-item').forEach(item => {
        item.classList.toggle('active', item.dataset.filter === filter);
    });

    // 清空搜索框
    searchInput.value = '';

    // 应用过滤
    applyFilter();
}

// 应用过滤
function applyFilter() {
    let filtered = portsData;

    switch (currentFilter) {
        case 'marked':
            filtered = portsData.filter(p => p.appMark);
            break;
        case 'docker':
            filtered = portsData.filter(p => p.container);
            break;
        case 'tcp':
            filtered = portsData.filter(p => p.protocol === 'tcp');
            break;
        case 'udp':
            filtered = portsData.filter(p => p.protocol === 'udp');
            break;
        default:
            filtered = portsData;
    }

    renderPorts(filtered);
}

// 搜索处理
function handleSearch(e) {
    const query = e.target.value.toLowerCase().trim();

    // 如果有搜索内容，重置过滤状态
    if (query) {
        currentFilter = 'all';
        document.querySelectorAll('.stat-item').forEach(item => {
            item.classList.remove('active');
        });
    }

    if (!query) {
        applyFilter();
        return;
    }

    const filtered = portsData.filter(port =>
        port.port.toString().includes(query) ||
        (port.processName && port.processName.toLowerCase().includes(query)) ||
        (port.appMark && port.appMark.appName.toLowerCase().includes(query)) ||
        (port.appMark && port.appMark.description && port.appMark.description.toLowerCase().includes(query))
    );

    renderPorts(filtered);
}

// ==================== 弹窗相关 ====================

// 打开编辑弹窗
function openEditModal(port, protocol, hasMark) {
    document.getElementById('editPort').value = port;
    document.getElementById('editProtocol').value = protocol;

    // 更新弹窗标题
    document.getElementById('modalTitle').textContent = hasMark ? t('editMark') : t('addMark');

    if (hasMark) {
        const portInfo = portsData.find(p => p.port === port && p.protocol === protocol);
        document.getElementById('appName').value = portInfo?.appMark?.appName || '';
        document.getElementById('description').value = portInfo?.appMark?.description || '';
    } else {
        document.getElementById('appName').value = '';
        document.getElementById('description').value = '';
    }

    modal.classList.add('active');
}

// 关闭弹窗
function closeModal() {
    modal.classList.remove('active');
}

// ==================== 标记操作 ====================

// 保存标记
async function handleSaveMark(e) {
    e.preventDefault();

    const data = {
        port: parseInt(document.getElementById('editPort').value),
        protocol: document.getElementById('editProtocol').value,
        appName: document.getElementById('appName').value.trim(),
        description: document.getElementById('description').value.trim()
    };

    try {
        const res = await fetch(`${API_BASE}/marks`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        });

        if (res.ok) {
            closeModal();
            showToast(t('markSaved'));

            // 更新本地数据，不重新扫描
            const portInfo = portsData.find(p => p.port === data.port && p.protocol === data.protocol);
            if (portInfo) {
                portInfo.appMark = {
                    appName: data.appName,
                    description: data.description,
                    port: data.port,
                    protocol: data.protocol
                };
            }

            // 刷新显示
            applyFilter();
            updateStats(portsData);
        } else {
            const err = await res.json();
            showToast(t('saveFailed') + (err.error ? ': ' + err.error : ''), true);
        }
    } catch (err) {
        showToast(t('saveFailed') + ': ' + err.message, true);
    }
}

// 删除标记
async function deleteMark(port, protocol) {
    if (!confirm(t('confirmDelete'))) return;

    try {
        const res = await fetch(`${API_BASE}/marks`, {
            method: 'DELETE',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ port, protocol })
        });

        if (res.ok) {
            showToast(t('markDeleted'));

            // 更新本地数据，不重新扫描
            const portInfo = portsData.find(p => p.port === port && p.protocol === protocol);
            if (portInfo) {
                portInfo.appMark = null;
            }

            // 刷新显示
            applyFilter();
            updateStats(portsData);
        } else {
            const err = await res.json();
            showToast(t('deleteFailed') + (err.error ? ': ' + err.error : ''), true);
        }
    } catch (err) {
        showToast(t('deleteFailed') + ': ' + err.message, true);
    }
}

// ==================== 进程终止 ====================

// 终止进程
async function killProcess(pid, processName) {
    if (!confirm(t('confirmKill') + `\nPID: ${pid}\nProcess: ${processName}`)) return;

    try {
        const res = await fetch(`${API_BASE}/kill/${pid}`, {
            method: 'POST'
        });

        if (res.ok) {
            showToast(t('processKilled'));
            // 进程终止后需要重新扫描
            setTimeout(() => {
                loadPorts();
            }, 500);
        } else {
            const err = await res.json();
            // 如果进程已经不存在，也算成功
            if (err.error && err.error.includes('no such process')) {
                showToast(t('processNotFound'));
                loadPorts();
            } else {
                showToast(t('killFailed') + ': ' + (err.error || 'Unknown error'), true);
            }
        }
    } catch (err) {
        showToast(t('killFailed') + ': ' + err.message, true);
    }
}

// ==================== Docker 容器操作 ====================

// Docker容器操作
async function dockerAction(containerId, action, containerName) {
    const confirmMessages = {
        stop: t('confirmStop'),
        start: t('confirmStart'),
        restart: t('confirmRestart'),
        remove: t('confirmRemove')
    };

    const successMessages = {
        stop: t('containerStopped'),
        start: t('containerStarted'),
        restart: t('containerRestarted'),
        remove: t('containerRemoved')
    };

    if (!confirm(confirmMessages[action] + `\n${t('container')}: ${containerName}`)) return;

    try {
        const res = await fetch(`${API_BASE}/docker/${containerId}/${action}`, {
            method: 'POST'
        });

        if (res.ok) {
            showToast(successMessages[action]);
            // 操作后需要重新加载数据
            setTimeout(() => {
                loadPorts();
                loadDockerStats();
            }, 500);
        } else {
            const err = await res.json();
            showToast(t('containerActionFailed') + ': ' + (err.error || 'Unknown error'), true);
        }
    } catch (err) {
        showToast(t('containerActionFailed') + ': ' + err.message, true);
    }
}

// 关闭容器详情弹窗
function closeContainerModal() {
    containerModal.classList.remove('active');
}

// ==================== 工具函数 ====================

// 显示提示
function showToast(message, isError = false) {
    const toast = document.getElementById('toast');
    toast.textContent = message;
    toast.className = 'toast show' + (isError ? ' error' : '');

    setTimeout(() => {
        toast.classList.remove('show');
    }, 3000);
}

// HTML 转义
function escapeHtml(str) {
    if (!str) return '';
    const div = document.createElement('div');
    div.textContent = str;
    return div.innerHTML;
}

// 点击弹窗外部关闭
modal.addEventListener('click', (e) => {
    if (e.target === modal) {
        closeModal();
    }
});

containerModal.addEventListener('click', (e) => {
    if (e.target === containerModal) {
        closeContainerModal();
    }
});

// ESC 关闭弹窗
document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape') {
        if (modal.classList.contains('active')) {
            closeModal();
        }
        if (containerModal.classList.contains('active')) {
            closeContainerModal();
        }
    }
});
