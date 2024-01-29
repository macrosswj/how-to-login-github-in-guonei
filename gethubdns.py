import requests

def gethubdns():
    hosts_path = r'C:\Windows\System32\drivers\etc\hosts'
    url = 'https://raw.hellogithub.com/hosts'
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36'
    }
    start_marker = "# GitHub520 Host Start"

    try:
        response = requests.get(url=url, headers=headers)
        response.raise_for_status()

        if response.status_code == 200:
            with open(hosts_path, 'r+', encoding='utf-8') as f:
                original_lines = f.readlines()
                insert_index = next((i for i, line in enumerate(original_lines) if start_marker in line), len(original_lines))

                # 清空从标记行到文件尾部的内容
                original_lines[insert_index:] = []

                # 添加新的内容
                original_lines[insert_index:insert_index] = [line + '\n' for line in response.text.splitlines()]

                # 写回文件
                f.seek(0)
                f.writelines(line for line in original_lines)
                f.truncate()
                print('写入成功')
        else:
            print('写入失败：响应状态码', response.status_code)
    except requests.exceptions.RequestException as e:
        print('请求异常:', e)
    except Exception as e:
        print('其他异常:', e)

# 调用函数
gethubdns()