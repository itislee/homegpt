# homegpt
家庭智能语音  
功能：向小爱音箱的提问时，自动转发至大语言模型并将模型的回答通过小爱音箱播放  

使用步骤

1. 设置环境变量  

export PROMPT_WORD="请问"  
export HOMEASSISTANT_IP="192.168.1.64"  
export HOMEASSISTANT_PORT="8123"  
export HOMEASSISTANT_TOKEN=""  
export WENXIN_KEYID=""  
export WENXIN_KEYSECRET=""  

2. 编译并运行程序  
sh build.sh
./homegpt

3. 构建Docker镜像  
docker build -t homegpt:v1 .
docker images
docker save {IMAGE_ID} -o homegpt.tar homegpt:v1